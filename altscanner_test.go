package altscanner_test

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/turtlemonvh/altscanner"
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

	ns := altscanner.NewAltScanner(shortLinesReader)
	for ns.Scan() {
		nlinesFound++
		assert.Equal(t, 300, len(ns.Bytes()))
		assert.Equal(t, 300, len(ns.Text()))
	}

	assert.Equal(t, nlines, nlinesFound)
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
	assert.True(t, ns.Err() == nil)
}

func TestLongLines(t *testing.T) {
	_, longLinesReader := generateReaders()

	nlinesFound := 0

	ns := altscanner.NewAltScanner(longLinesReader)
	for ns.Scan() {
		nlinesFound++
		assert.Equal(t, ncharslonglines, len(ns.Text()))
	}

	assert.Equal(t, nlines, nlinesFound)
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
	assert.True(t, ns.Err() != nil)
}
