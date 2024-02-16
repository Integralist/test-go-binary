# test-go-binary

With no git tags:

> **NOTE:** Unless mentioned, the `go.mod` module name is `github.com/integralist/test-go-binary`.

```shell
$ go install github.com/integralist/test-go-binary@latest

go: downloading github.com/integralist/test-go-binary v0.0.0-20240216133252-d29459c6bbac

$ file ~/go/bin/test-go-binary

/Users/integralist/go/bin/test-go-binary: Mach-O 64-bit executable arm64
```

With a single v1 git tag:

```shell
$ go install github.com/integralist/test-go-binary@v1

go: github.com/integralist/test-go-binary@v1: no matching versions for query "v1"

$ go install github.com/integralist/test-go-binary@v1.0.0

go: github.com/integralist/test-go-binary@v1.0.0: github.com/integralist/test-go-binary@v1.0.0: invalid version: unknown revision v1.0.0
```

With a v1.0.0 git tag:

```shell
$ go install github.com/integralist/test-go-binary@latest

go: downloading github.com/integralist/test-go-binary v1.0.0
```

> **NOTE:** Wait for [pkg.go.dev/github.com/integralist/test-go-binary](https://pkg.go.dev/github.com/integralist/test-go-binary) (here's the [v1.0.0](https://pkg.go.dev/github.com/integralist/test-go-binary@v1.0.0) which takes a while to be published).

With a v2.0.0 git tag:

```shell
$ go install github.com/integralist/test-go-binary@latest

# no output

$ go install github.com/integralist/test-go-binary@v2.0.0

go: github.com/integralist/test-go-binary@v2.0.0: github.com/integralist/test-go-binary@v2.0.0: invalid version: module contains a go.mod file, so module path must match major version ("github.com/integralist/test-go-binary/v2")
```

It seems to install the previous version, which was the v1.0.0 tag.

But when explicitly requesting v2.0.0 it tells me the go module name should be
updated to reflect the major version change (as v1.x.x doesn't require a v1 in
the module name). This error would also indicate why (no matter how long you
wait) the [test-go-binary@v2.0.0](https://pkg.go.dev/github.com/integralist/test-go-binary@v2.0.0)
never works (you need a `test-go-binary/v2` path instead).

With a v2.0.0 git tag + go module name now including `/v2`:

> **NOTE:** You need [pkg.go.dev/github.com/integralist/test-go-binary/v2](https://pkg.go.dev/github.com/integralist/test-go-binary/v2)

```shell
$ go install github.com/integralist/test-go-binary/v2@latest
```

## Reference

- [v2-go-modules](https://go.dev/blog/v2-go-modules)
- [cmd/go](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)
