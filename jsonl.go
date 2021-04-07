package jsonl

import (
	"bufio"
	"encoding/json"
	"io"
)

type Decoder struct {
	scanner *bufio.Scanner
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{scanner: bufio.NewScanner(r)}
}

func (dec *Decoder) More() bool {
	return dec.scanner.Scan()
}

func (dec *Decoder) Decode(v interface{}) error {
	if dec.scanner.Err() != nil {
		return dec.scanner.Err()
	}
	return json.Unmarshal(dec.scanner.Bytes(), &v)
}
