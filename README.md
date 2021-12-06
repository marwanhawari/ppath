<p align="center">
  <img width=70% height=auto src="https://github.com/marwanhawari/ppath/raw/main/assets/ppath_logo.png" alt="ppath logo"/>
</p>

[![Build Status](https://github.com/marwanhawari/ppath/actions/workflows/build.yml/badge.svg)](https://github.com/marwanhawari/ppath/actions)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](https://github.com/marwanhawari/ppath/blob/main/CODE_OF_CONDUCT.md)
[![GitHub](https://img.shields.io/github/license/marwanhawari/ppath?color=blue)](https://github.com/marwanhawari/ppath/blob/main/LICENSE)

# Description
A command line tool to pretty print your system's PATH variable. The output paths are colorized if they have special associations (e.g. cyan for go, green for python, orange for rust).

# Installation
Install directly with go:
```
go install github.com/marwanhawari/ppath@latest
```

# Usage


### Options
```
Usage of ppath:
  -all
        Include duplicate paths from $PATH variable. (default false)
  -colorize string
        Choose how to color the output [8bit, 3bit, none]. (default "8bit")
```