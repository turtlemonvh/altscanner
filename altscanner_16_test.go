//+build go1.6

package altscanner_test

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/turtlemonvh/altscanner"
	"strings"
	"testing"
)

type Bufferable interface {
	Buffer([]byte, int)
}

// Only works on go>1.6, because that's when the `Scanner.Buffer` method was added
// https://golang.org/doc/go1.6#minor_library_changes
func TestLongLinesDefaultScannerLargeBytes(t *testing.T) {
	_, longLinesReader := generateReaders()

	nlinesFound := 0

	// Should pass
	ns := bufio.NewScanner(longLinesReader)
	ns.Buffer(make([]byte, bufio.MaxScanTokenSize), bufio.MaxScanTokenSize*10)
	for ns.Scan() {
		nlinesFound++
		assert.Equal(t, ncharslonglines, len(ns.Text()))
	}

	assert.Equal(t, 5, nlinesFound)
	assert.True(t, ns.Err() == nil)
}

func benchBufioScanner(b *testing.B, content string) {
	for n := 0; n < b.N; n++ {
		scanner := bufio.NewScanner(strings.NewReader(content))
		for scanner.Scan() {
			scanner.Bytes()
		}
	}
}

func benchBufferedBufioScanner(b *testing.B, content string) {
	for n := 0; n < b.N; n++ {
		scanner := bufio.NewScanner(strings.NewReader(content))
		scanner.Buffer(make([]byte, bufio.MaxScanTokenSize), bufio.MaxScanTokenSize*10)
		for scanner.Scan() {
			scanner.Bytes()
		}
	}
}

func benchAltScanner(b *testing.B, content string) {
	for n := 0; n < b.N; n++ {
		scanner := altscanner.NewAltScanner(strings.NewReader(content))
		for scanner.Scan() {
			scanner.Bytes()
		}
	}
}

func BenchmarkBufioScannerSmall(b *testing.B)         { benchBufioScanner(b, shortLines) }
func BenchmarkBufferedBufioScannerSmall(b *testing.B) { benchBufioScanner(b, shortLines) }
func BenchmarkAltScannerSmall(b *testing.B)           { benchAltScanner(b, shortLines) }
func BenchmarkBufferedBufioScannerLong(b *testing.B)  { benchBufioScanner(b, longLines) }
func BenchmarkAltScannerLong(b *testing.B)            { benchAltScanner(b, longLines) }
