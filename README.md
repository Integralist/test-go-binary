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

> **NOTE:** Wait for [pkg.go.dev/github.com/integralist/test-go-binary](https://pkg.go.dev/github.com/integralist/test-go-binary)

## Reference

- [cmd/go](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)
