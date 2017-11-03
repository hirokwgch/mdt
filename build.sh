# available $GOOS and $GOARCH
# https://golang.org/doc/install/source#environment
MDT_VERSION=1.0.0

mkdir archive 2> /dev/null

GOOS=windows GOARCH=amd64 go build -o "archive/mdt-${MDT_VERSION}-win-amd64.exe"
GOOS=linux GOARCH=amd64 go build -o "archive/mdt-${MDT_VERSION}-linux-amd64"