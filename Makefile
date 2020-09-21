sync:
	rsync -aP gobot -i gobot aleksandrhorkavyi@35.232.133.126:/home/aleksandrhorkavyi/work/src/github.com/aleksandrhorkavyi/gobot/

build-ubuntu:
	GOOS=linux GOARCH=amd64 go build -v .