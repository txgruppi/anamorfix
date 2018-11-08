# anamorfix

![License](https://img.shields.io/github/license/txgruppi/anamorfix.svg?style=flat-square)
[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg?longCache=true&style=flat-square)](https://godoc.org/github.com/txgruppi/anamorfix)
[![GitHub release](https://img.shields.io/github/release/txgruppi/anamorfix.svg?style=flat-square)](https://github.com/txgruppi/anamorfix/releases)

_Copy_ of https://github.com/shazbits/anamorfix, but in Go.

## Why Go?

Single binary, no dependencies.

## The command

### Install

You can download the command from the [Releases page](https://github.com/txgruppi/anamorfix/releases) or, if you have Go in your machine, you can install by running

```sh
go get -u github.com/txgruppi/anamorfix/cmd/anamorfix
```

### Usage

```sh
anamorfix ./IMG_*
```

Easy, right?

Check out the available options:

```plain
NAME:
   anamorfix - Fix aspect ratio of anamorphic images

USAGE:
   anamorfix [global options] command [command options] <files>

VERSION:
   0.0.0

AUTHOR:
   Tarcisio Gruppi <txgruppi@gmail.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --width, -w               Resize image's width (default: resize image's height)
   --ratio value, -r value   resize ratio (default: 1.33)
   --suffix value, -s value  name suffix to add to resized files (default: "-desqueezed")
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```

## The library

You can use this code in your project.

### Install

You can download the library by running

```sh
go get -u github.com/txgruppi/anamorfix
```

### Documentation

Check the docs at [GoDoc](https://godoc.org/github.com/txgruppi/anamorfix).
