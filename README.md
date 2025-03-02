# Lab1: Quadratic Equation Solver
- [Overview](#overview)
- [Installation and Usage](#installation-and-usage)
  - [via Docker](#via-docker)
  - [via Go](#via-go)
- [File format for non-interactive mode](#file-format-for-non-interactive-mode)
- [Revert commit](#revert-commit)
## Overview

This application solves quadratic equations of the form `ax² + bx + c = 0`. It provides two modes of operation:  

- **Interactive mode**: The user enters coefficients step by step.  
- **Non-interactive mode**: The coefficients are read from a file provided as a command-line argument.  

The program is a console application designed for quick and efficient calculations.  

## Installation and Usage

### via Docker

To run in interactive mode:
```sh
docker run --rm -i rmerezha/quadratic-solver:1.0.1
```

or

```sh
docker build -t quadratic-solver:latest .
docker run --rm -i quadratic-solver:latest
```

To run in non-interactive mode:

```sh
docker run --rm -v /path/to/host/file:/path/to/container/file \
                rmerezha/quadratic-solver:1.0.1 /path/to/container/file
```

or

```sh
docker build -t quadratic-solver:latest .
docker run --rm -v /path/to/host/file:/path/to/container/file \
                quadratic-solver:latest /path/to/container/file
```

### via Go

**Go version 1.23.5 or higher**

To run in interactive mode:

```sh
go build -o solver
./solver
```

To run in non-interactive mode:

```sh
go build -o solver
./solver /path/to/file
```

## File format for non-interactive mode

File format:

```
a b c
```

For example:

```
1.23 4.5678 9
```

```
1.23<-space->4.5678<-space->9<-spaces/new lines/tabs->
```

## Revert commit
[b6bd64b45ff545aed8bb022d3a6dd9328b661c23](https://github.com/rmerezha/mtrpz-lab1/commit/b6bd64b45ff545aed8bb022d3a6dd9328b661c23)
