@echo off
echo "Calling the Hasselhoff..."
SET VERSION=0.2.1

curl -s -L https://github.com/angelbarrera92/hasselhoffme/releases/download/%VERSION%/hasselhoffme_%VERSION%_windows_amd64.exe --output hass.exe > NUL

echo "Hasselhoffing..."
hass.exe
