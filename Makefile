VERSION=`git describe --tags --dirty --always`

LDFLAGS=-ldflags "-X main.version=${VERSION}"

build:
	go build ${LDFLAGS} ${OPTS}

generate:
	go generate github.com/regiohelden/innovazammad/innovaphone