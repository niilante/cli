arukas/cli
==========

* Website: https://arukas.io

# Arukas Cloud Command Line Interface

* Get api key here https://app.arukas.io/settings/api-keys
* export it

```bash
export ARUKAS_JSON_API_SECRET=hoge
export ARUKAS_JSON_API_TOKEN=hoge
```

### Cross Compilation and Building for Distribution

If you wish to cross-compile arukas-cli for another architecture, you can set the `XC_OS` and `XC_ARCH` environment variables to values representing the target operating system and architecture before calling `make`. The output is placed in the `pkg` subdirectory tree both expanded in a directory representing the OS/architecture combination and as a ZIP archive.

For example, to compile 64-bit Linux binaries on Mac OS X Linux, you can run:

```sh
$ XC_OS=linux XC_ARCH=amd64 make bin
...
$ file pkg/linux_amd64/arukas
arukas: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
```

`XC_OS` and `XC_ARCH` can be space separated lists representing different combinations of operating system and architecture. For example, to compile for both Linux and Mac OS X, targeting both 32- and 64-bit architectures, you can run:

```sh
$ XC_OS="linux darwin" XC_ARCH="386 amd64" make bin
...
$ tree ./pkg/ -P "arukas|*.zip"
./pkg/
├── darwin_386
│   └── arukas
├── darwin_386.zip
├── darwin_amd64
│   └── arukas
├── darwin_amd64.zip
├── linux_386
│   └── arukas
├── linux_386.zip
├── linux_amd64
│   └── arukas
└── linux_amd64.zip

4 directories, 8 files
```

_Note: Cross-compilation uses [gox](https://github.com/mitchellh/gox), which requires toolchains to be built with versions of Go prior to 1.5. In order to successfully cross-compile with older versions of Go, you will need to run `gox -build-toolchain` before running the commands detailed above._
