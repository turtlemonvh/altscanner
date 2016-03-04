package altscanner

import (
	"bufio"
	"io"
)

// https://golang.org/src/bufio/scan.go?s=5433:5462#L118

// Alternative to scanner that works with lines longer than the buffer size
// Take an io.Reader, return a function that you can loop over

type AltScanner struct {
	reader      *bufio.Reader
	currentLine string
	err         error
}

func NewAltScanner(r io.Reader) *AltScanner {
	return &AltScanner{
		reader:      bufio.NewReader(r),
		currentLine: "",
		err:         nil,
	}
}

func (s *AltScanner) Text() string {
	return s.currentLine
}

func (s *AltScanner) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}

func (s *AltScanner) Scan() bool {
	partialLine := []byte{}
	s.currentLine = ""
	prefix := true

	for true {
		partialLine, prefix, s.err = s.reader.ReadLine()
		if s.err != nil {
			return false
		}

		// Add this component of the line
		s.currentLine += string(partialLine)

		if !prefix {
			// Finished reading this line
			return true
		}
	}

	return false
}
