VERSION=$(shell git describe --long --tags)
DATE=$(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS="-X main.date=$(DATE) -X main.version=$(VERSION)"

all: pblog

templates_embed.go: templates/*.html
	go generate

pblog: pblog.go templates_embed.go
	go build -ldflags $(LDFLAGS) pblog.go templates_embed.go

pblog.linux: pblog.go templates_embed.go
	env GOOS=linux GOARCH=amd64 go build -o pblog.linux -ldflags $(LDFLAGS) pblog.go templates_embed.go

cgi-server: cgi-server.go
	go build cgi-server.go

deploy-prod: pblog.linux
	./post-v1.sh

run: pblog
	./pblog

run-prod:
	../aceapi/apiexec ./pblog.sh
	../aceapi/apiexec tail -50 ../../pblog/pblog_err.log

clean:
	rm -f pblog cgi-server pblog.linux
