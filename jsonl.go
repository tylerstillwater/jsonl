package jsonl

import (
	"bufio"
	"encoding/json"
	"io"
)

type Decoder struct {
	scanner   *bufio.Scanner
	debug     bool
	debugData []string
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{scanner: bufio.NewScanner(r)}
}

// SetDebug sets the internal debug flag to collect debug data. This data can be
// printed with the PrintDebug() method.
func (dec *Decoder) SetDebug(debug bool) {
	if !debug {
		dec.debugData = nil
	}
	dec.debug = debug
}

// PrintDebug iterates the debugData slice and prints each line. Once printed, it
// resets the slice.
func (dec *Decoder) PrintDebug() {
	for _, line := range dec.debugData {
		println(line)
	}
	dec.clearDebug()
}

func (dec *Decoder) More() bool {
	out := dec.scanner.Scan()
	if dec.debug {
		dec.debugData = append(dec.debugData, dec.scanner.Text())
	}
	return out
}

// Decode decodes the next JSON object from the input stream. If the operation succeeds,
// the debug slice is reset.
func (dec *Decoder) Decode(v interface{}) error {
	if dec.scanner.Err() != nil {
		return dec.scanner.Err()
	}
	err := json.Unmarshal(dec.scanner.Bytes(), &v)
	if err != nil {
		return err
	}
	dec.clearDebug()
	return nil
}

// clearDebug clears the debug slice.
func (dec *Decoder) clearDebug() {
	if dec.debug {
		dec.debugData = dec.debugData[:0]
	}
}
