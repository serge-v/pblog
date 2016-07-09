VERSION=$(shell git describe --long --tags)
DATE=$(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS="-X main.date=$(DATE) -X main.version=$(VERSION)"

all: pblog cgi-server

pblog: pblog.go
	go build -ldflags $(LDFLAGS) pblog.go

pblog.linux: pblog.go
	env GOOS=linux GOARCH=amd64 go build -o pblog.linux -ldflags $(LDFLAGS) pblog.go

cgi-server: cgi-server.go
	go build cgi-server.go

deploy-prod: pblog.linux
