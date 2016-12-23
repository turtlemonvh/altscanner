//+build go1.6

package altscanner

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
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
	fmt.Println("ns.Err()", ns.Err())
	assert.True(t, ns.Err() == nil)
}
