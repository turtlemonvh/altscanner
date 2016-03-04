# AltScanner [![GoDoc](https://godoc.org/github.com/turtlemonvh/altscanner?status.svg)](https://godoc.org/github.com/turtlemonvh/altscanner) [![Build Status](https://travis-ci.org/turtlemonvh/altscanner.png?branch=master)](https://travis-ci.org/turtlemonvh/altscanner)

A version of `bufio.Scanner` that works with lines of arbitrary length.

## Why

If you're getting a `bufio.Scanner: token too long` error, this may be what you want.

## Caveats

* Only the `.Text()` method of the scanner works. 
* Only breaks on newlines. 

## Alternatives

If you have a good idea about the size of your data and are running go>1.6 ([where the `Scanner.Buffer` method was introduced](https://golang.org/doc/go1.6#minor_library_changes)), you probably just want to change the size of the buffer used by the scanner.  For example:

    // Create a scanner and resize its buffer to be 10X larger than usual (640 Kb instead of 64 Kb)
    scanner := bufio.NewScanner(file)
    scanner.Buffer(make([]byte, bufio.MaxScanTokenSize), bufio.MaxScanTokenSize*10)

However, if you need to be compatible with go<1.6 or you really have no idea about the size of your data, this approach works pretty well.

## License

MIT
