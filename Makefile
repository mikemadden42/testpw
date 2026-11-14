BINARY := testpw
LDFLAGS := -s -w
PLATFORMS := darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm64 windows/386 windows/amd64 windows/arm64

build:
	@set -e; for p in $(PLATFORMS); do \
		os=$${p%/*}; arch=$${p#*/}; ext=; \
		[ "$$os" = windows ] && ext=.exe; \
		echo "building $$os/$$arch"; \
		GOOS=$$os GOARCH=$$arch go build -ldflags='$(LDFLAGS)' -o ./bin/$(BINARY)_$${os}_$${arch}$$ext .; \
	done

test:
	go test ./...

check:
	golangci-lint run
	gosec ./...

release:
	goreleaser release --snapshot --clean

clean:
	rm -rf bin dist

.PHONY: build test check release clean
