<p align="center">
  <img width=70% height=auto src="https://github.com/marwanhawari/ppath/raw/main/assets/ppath_logo.png" alt="ppath logo"/>
</p>

<p align="center">
  <a href="https://github.com/marwanhawari/ppath/actions/">
    <img src="https://github.com/marwanhawari/ppath/actions/workflows/test.yml/badge.svg" alt="build status"/>
  </a>

  <a href="https://goreportcard.com/report/github.com/marwanhawari/ppath">
    <img src="https://goreportcard.com/badge/github.com/marwanhawari/ppath" alt="go report card"/>
  </a>

  <a href="https://github.com/marwanhawari/ppath/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/marwanhawari/ppath?color=blue" alt="license"/>
  </a>
</p>

# Description
A command-line tool to pretty print your system's PATH environment variable. The output paths are colorized if they have special associations (e.g. cyan for go, green for python, orange for rust).

# Installation
Install a pre-compiled binary:
```
brew install marwanhawari/tap/ppath
```

Install from source:
```
go install github.com/marwanhawari/ppath@latest
```

# Usage
![ppath usage](https://github.com/marwanhawari/ppath/raw/main/assets/ppath_usage.gif)

### Options
```
Usage of ppath:
  -all
        Include duplicate paths from $PATH variable. (default false)
  -uncolored
        Uncolor the output. (default false)
```
