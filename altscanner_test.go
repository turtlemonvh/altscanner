package altscanner

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

var (
	shortLines      string
	longLines       string
	nlines          = 5
	ncharslonglines = 300000
)

// Create fixtures
func init() {
	linestarter := "frog mule horse cat dog mouse " // 30

	// Long lines
	// Need longer than 64 KB lines
	// https://golang.org/src/bufio/scan.go#L71
	longLines = ""
	for i := 0; i < nlines; i++ {
		// 300K long lines
		for j := 0; j < 10000; j++ {
			longLines += linestarter
		}
		longLines += "\n"
	}

	// Short lines
	shortLines = ""
	for i := 0; i < nlines; i++ {
		for j := 0; j < 10; j++ {
			shortLines += linestarter
		}
		shortLines += "\n"
	}
}

// Create some test io.Readers
// Needs to be called for each test
func generateReaders() (io.Reader, io.Reader) {
	return strings.NewReader(shortLines), strings.NewReader(longLines)
}

func TestShortLines(t *testing.T) {
	shortLinesReader, _ := generateReaders()

	nlinesFound := 0

	ns := NewAltScanner(shortLinesReader)
	for ns.Scan() {
		nlinesFound++
		assert.Equal(t, 300, len(ns.Bytes()))
		assert.Equal(t, 300, len(ns.Text()))
	}

	assert.Equal(t, nlines, nlinesFound)
	fmt.Println("ns.Err()", ns.Err())
	assert.True(t, ns.Err() == nil)
}

func TestShortLinesDefaultScanner(t *testing.T) {
	shortLinesReader, _ := generateReaders()

	nlinesFound := 0

	ns := bufio.NewScanner(shortLinesReader)
	for ns.Scan() {
		nlinesFound++
	}

	assert.Equal(t, nlines, nlinesFound)
	fmt.Println("ns.Err()", ns.Err())
	assert.True(t, ns.Err() == nil)
}

func TestLongLines(t *testing.T) {
	_, longLinesReader := generateReaders()

	nlinesFound := 0

	ns := NewAltScanner(longLinesReader)
	for ns.Scan() {
		nlinesFound++
		assert.Equal(t, ncharslonglines, len(ns.Text()))
	}

	assert.Equal(t, nlines, nlinesFound)
	fmt.Println("ns.Err()", ns.Err())
	assert.True(t, ns.Err() == nil)
}

func TestLongLinesDefaultScanner(t *testing.T) {
	_, longLinesReader := generateReaders()

	nlinesFound := 0

	ns := bufio.NewScanner(longLinesReader)
	for ns.Scan() {
		assert.True(t, ns.Err() != nil)
		nlinesFound++
	}

	assert.Equal(t, 0, nlinesFound)
	fmt.Println("ns.Err()", ns.Err())
	assert.True(t, ns.Err() != nil)
}

/*

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

*/
