# testpw

Go tool that tests password strength using [zxcvbn](https://github.com/ccojocar/zxcvbn-go).

## Install

```sh
go install github.com/mikemadden42/testpw@latest
```

## Usage

Run it interactively. Input is hidden while you type; press Ctrl-C to exit.

```sh
$ testpw
Enter password or Ctrl-c to exit:

Password score    (0-4): 4
Estimated entropy (bit): 43.19
Estimated time to crack: 2.0 years
```

Or pipe passwords in, one per line:

```sh
testpw < passwords.txt
```

## Development

```sh
make test     # run tests
make check    # golangci-lint (incl. gofmt/goimports) and gosec
make build    # cross-compile into ./bin
make release  # goreleaser snapshot into ./dist
```
