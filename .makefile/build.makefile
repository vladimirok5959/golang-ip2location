BINARIES=$(patsubst cmd/%/main.go,bin/%,$(wildcard cmd/*/main.go))
STATIC_BINARIES=$(patsubst cmd/%/main.go,bin/%-static,$(wildcard cmd/*/main.go))

all: $(BINARIES)
static: $(STATIC_BINARIES)

.PHONY: all static

GO_SOURCES=
GO_SOURCES+=$(shell find . -name '*.go' -not -path '*vendor*' -not -path '*.pb.go')
GO_SOURCES+=$(wildcard go.*)

bin/%: cmd/%/main.go $(GO_SOURCES)
	@mkdir -p $(dir $@)
	go build -o $@ $<

bin/%-static: cmd/%/main.go $(GO_SOURCES)
	@mkdir -p $(dir $@)
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o $@ $<
