GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

TAGS=mattn

cli:
	go build -tags $(TAGS) -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/wof-sqlite-index cmd/wof-sqlite-index/main.go
