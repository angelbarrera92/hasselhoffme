package motd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/zyxar/image2ascii/ascii"
)

const (
	// MOTDFile is a path to /etc/motd file
	MOTDFile = "/etc/motd"
	// UpdateMOTDPath is a path to /etc/update-motd.d
	UpdateMOTDPath = "/etc/update-motd.d"
	// UpdateMOTDFile is a path to /etc/update-motd.d/99-hasselhoffme
	UpdateMOTDFile = UpdateMOTDPath + "/99-hasselhoffme"
)

func SetMotd(data []byte) {
	if os.Getuid() != 0 {
		fmt.Fprintf(os.Stderr, "motd can only be changed as root\n")
		os.Exit(1)
	}

	opt := ascii.Options{
		Width:  0,
		Height: 0,
		Color:  false,
		Invert: false,
		Flipx:  false,
		Flipy:  false}
	motd, err := ascii.Decode(bytes.NewReader(data), opt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while decoding data: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat(UpdateMOTDPath); os.IsNotExist(err) {
		writeMotd(motd)
	} else {
		writeUpdateMotdScript(motd)
	}
}

func writeMotd(motd *ascii.Ascii) {
	content := ""
	if _, err := os.Stat(MOTDFile); !os.IsNotExist(err) {
		contentBytes, err := ioutil.ReadFile(MOTDFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while opening %s: %v\n", MOTDFile, err)
			os.Exit(1)
		}
		content = string(contentBytes)
	}

	re := regexp.MustCompile(`(?s)### hasselhon ###.*### hasselhoff ###\n`)
	content = re.ReplaceAllString(content, "")

	f, err := os.OpenFile(MOTDFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755) // nolint
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening %s: %v\n", MOTDFile, err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s### hasselhon ###\n", content)
	_, err = motd.WriteTo(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while writing to file %s: %v\n", f.Name(), err)
		os.Exit(1)
	}
	fmt.Fprintf(f, "### hasselhoff ###\n")
}

func writeUpdateMotdScript(motd *ascii.Ascii) {
	f, err := os.OpenFile(UpdateMOTDFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755) // nolint
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while opening %s: %v\n", UpdateMOTDFile, err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, `#!/bin/sh
cat <<EOF
`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while outputting to file %s: %v\n", f.Name(), err)
		os.Exit(1)
	}

	_, err = motd.WriteTo(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while writing to file %s: %v\n", f.Name(), err)
		os.Exit(1)
	}

	_, err = fmt.Fprintf(f, "EOF\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while outputting to file %s: %v\n", f.Name(), err)
		os.Exit(1)
	}
}
