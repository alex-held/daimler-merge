# daimler-merge

coding task 2 for a job interview

[![Go Report Card](https://goreportcard.com/badge/github.com/alex-held/daimler-merge)](https://goreportcard.com/report/github.com/alex-held/daimler-merge)
[![codecov](https://codecov.io/gh/alex-held/daimler-merge/branch/master/graph/badge.svg?token=xlCvdqX45x)](https://codecov.io/gh/alex-held/daimler-merge)
[![Coverage Status](https://img.shields.io/codecov/c/github/alex-held/daimler-merge.svg)](https://codecov.io/gh/alex-held/daimler-merge)
[![Release](https://github.com/alex-held/daimler-merge/workflows/Release/badge.svg)](https://github.com/alex-held/daimler-merge/releases)
[![wakatime](https://wakatime.com/badge/github/alex-held/daimler-merge.svg)](https://wakatime.com/badge/github/alex-held/daimler-merge)

## Overview

**This project gets not actively maintained!**

This project provides an algorithm to merge overlapping integer intervals. You can find this
algorithm [inside the merge package](pkg/merge).

Refer to [Task](docs/Task.adoc) for detailed information.

For a report on time and space complexity please visit [Complexity](docs/Complexity.adoc). 

## Build

## Makefile

``` shell
# build for multi-platform
make build-any

# build for current platform
make build
```

## Docker
``` shell
# build
docker compose build

# run
docker compose run daimler-merge --version
```

## Install

``` shell
go get github.com/alex-held/daimler-merge
```

## Usage

``` shell
daimler-merge
// Before: [[2 19] [4 8] [14 23] [25 30]]
// After: [[2 23] [25 30]]
```

``` shell
daimler-merge --help
daimler-merge --version
```
