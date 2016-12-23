# AltScanner [![GoDoc](https://godoc.org/github.com/turtlemonvh/altscanner?status.svg)](https://godoc.org/github.com/turtlemonvh/altscanner) [![Build Status](https://travis-ci.org/turtlemonvh/altscanner.png?branch=master)](https://travis-ci.org/turtlemonvh/altscanner)

A version of `bufio.Scanner` that works with lines of arbitrary length.

## Why

If you're getting a `bufio.Scanner: token too long` error, this may be what you want.

## How

If your code used to look like this:

```golang
import "bufio"

s := bufio.NewScanner(myIoReader)
for s.Scan() {
    // Do work
}
```

You can now handle very long lines without errors by changing to:

```golang
import "github.com/turtlemonvh/altscanner"

s := altscanner.NewAltScanner(myIoReader)
for s.Scan() {
    // Do work
}
```

## Caveats

* Only breaks on newlines.
* Just appends bytes to a byte slice instead of using [a real buffer](https://golang.org/pkg/bytes/#Buffer).

## Alternatives

If you have a good idea about the size of your data and are running go>1.6 ([where the `Scanner.Buffer` method was introduced](https://golang.org/doc/go1.6#minor_library_changes)), you probably just want to change the size of the buffer used by the scanner.  For example:

    // Create a scanner and resize its buffer to be 10X larger than usual (640 Kb instead of 64 Kb)
    scanner := bufio.NewScanner(file)
    scanner.Buffer(make([]byte, bufio.MaxScanTokenSize), bufio.MaxScanTokenSize*10)

However, if you need to be compatible with go<1.6 or you really have no idea about the size of your data, this approach works pretty well.

## Performance

It is robust, but not very fast.  The benchmark results below show the performance of reading in 5 lines of content. The lines used in the tests are either 30 bytes (short) or 300K bytes (long).

```bash
$ go test -test.bench=Scanner -test.run=^$ -test.benchmem
BenchmarkBufioScannerSmall-8             1000000          1061 ns/op        4128 B/op          2 allocs/op
BenchmarkBufferedBufioScannerSmall-8     1000000          1059 ns/op        4128 B/op          2 allocs/op
BenchmarkAltScannerSmall-8               1000000          1779 ns/op        5824 B/op          8 allocs/op
BenchmarkBufferedBufioScannerLong-8        50000         28077 ns/op      127008 B/op          6 allocs/op
BenchmarkAltScannerLong-8                   2000       1142195 ns/op     7032704 B/op         78 allocs/op
PASS
ok      github.com/turtlemonvh/altscanner   13.458s
```

`AltScanner` is significantly slower, has many more allocations, and uses significantly more bytes per operation than the buffer `bufio.Scanner`.  In short: it is always faster to use `Scanner.Buffer` to adjust the size of the buffer if you are using go1.6+ and you are confident about the max possible size of an line.

## License

MIT
