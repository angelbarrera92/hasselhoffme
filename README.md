```                                                
   __ __   ___    ____   ____   ____   __    __ __  ____    ____   ____   __  ___   ____
  / // /  / _ |  / __/  / __/  / __/  / /   / // / / __ \  / __/  / __/  /  |/  /  / __/
 / _  /  / __ | _\ \   _\ \   / _/   / /__ / _  / / /_/ / / _/   / _/   / /|_/ /  / _/  
/_//_/  /_/ |_|/___/  /___/  /___/  /____//_//_/  \____/ /_/    /_/    /_/  /_/  /___/                                                                                   
```

![Header](.readme_images/header.png?raw=true "Header")

[![Release](https://img.shields.io/github/release/angelbarrera92/hasselhoffme.svg)](https://github.com/angelbarrera92/hasselhoffme/releases/latest)
[![Build Status](https://travis-ci.org/angelbarrera92/hasselhoffme.svg?branch=master)](https://travis-ci.org/angelbarrera92/hasselhoffme)
[![Go Report Card](https://goreportcard.com/badge/github.com/angelbarrera92/hasselhoffme)](https://goreportcard.com/badge/github.com/angelbarrera92/hasselhoffme)
[![Github All Releases](https://img.shields.io/github/downloads/angelbarrera92/hasselhoffme/total.svg)](https://www.somsubhra.com/github-release-stats/?username=angelbarrera92&repository=hasselhoffme)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg)](https://github.com/goreleaser)

## Why does this even exist?
Security. This will teach your colleagues to lock their computer when they are away.

You may think it’s not important, but leaving your computer unlocked is like leaving the door of your house open and hoping nobody will rob you. Anyone could sit at your computer and gain access to your private information.

Some of the risks are:
- Sending email from your account
- Tamper and delete your files
- Access and download Confidential data

## How to use?

Open the terminal and type the following command:

### Wget
```bash
wget https://tiny.cc/hasselhoff -O - | bash
```

### Curl
```bash
curl -L https://tiny.cc/hasselhoff | bash
```

If you can't run the script, try with this url:
```bash
curl -L https://tiny.cc/hasselhoffme | bash
```

## Windows usage

```cmd
curl -s -L https://tiny.cc/hasselhoffme-win | cmd
```

or 


```cmd
curl -s -L https://tiny.cc/hasselhoff-win | cmd
```

## Tested
- Ubuntu 16 and 18
- Fedora 27 and 28
- MacOS
- Windows 10

## Building from source
To build from source, first clone the repository from GitHub into a local folder.

```bash
git clone https://github.com/angelbarrera92/hasselhoffme.git
```

Change folders into the newly clones working copy and ensure the project dependancies are available
```bash
export GO111MODULE=on
go mod download
```

Next build
```bash
go build -v
```

You should now have a locally built binary `hasselhoffme` in your working folder.

## Stargazers over time
[![Stargazers over time](https://starcharts.herokuapp.com/angelbarrera92/hasselhoffme.svg)](https://starcharts.herokuapp.com/angelbarrera92/hasselhoffme)

## Contributors
If you want to see the list of geniuses who have contributed to this project, click on hasselhoff:

[![Hoff](https://media.giphy.com/media/UuDxS2EBRZSyA/giphy.gif)](./CONTRIBUTORS.md)
