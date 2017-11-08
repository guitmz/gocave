[![Build Status](https://travis-ci.org/guitmz/gocave.svg?branch=master)](https://travis-ci.org/guitmz/gocave)

[![](https://images.microbadger.com/badges/image/guitmz/gocave.svg)](https://microbadger.com/images/guitmz/gocave "Get your own image badge on microbadger.com")


# gocave

Utility to find code caves in ELF files, written in Go.

# Installation

You can either run `go get -u github.com/guitmz/gocave` or clone this repository and build with `go build`.

# Usage
`$ gocave elf_file cave_size`

Or with Docker:  

```
$ docker pull guitmz/gocave
$ docker run --rm -v elf_file:/elf_file guitmz/gocave /elf_file cave_size
```

# Todo
- Add payload injection and execution;
- Add more binary types like Mach and PE.
- ?

# References
https://www.codeproject.com/Articles/20240/The-Beginners-Guide-to-Codecaves

https://github.com/Antonin-Deniau/cave_miner
