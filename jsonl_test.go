package jsonl

import (
	"strings"
	"testing"

	"github.com/tylerstillwater/proof"
)

func TestDecoder(t *testing.T) {
	t.Parallel()
	prove := proof.New(t)

	const input = `{"name":"first", "value":1}
{"name":"second", "value":2}`

	type testRow struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
	var rows []testRow
	decoder := NewDecoder(strings.NewReader(input))
	for decoder.More() {
		var row testRow
		prove.NotErr(decoder.Decode(&row))
		rows = append(rows, row)
	}

	prove.Len(rows, 2)
	prove.Lax(
		func(lax *proof.Prover) {
			lax.Equal(rows[0].Name, "first")
			lax.Equal(rows[0].Value, 1)
			lax.Equal(rows[1].Name, "second")
			lax.Equal(rows[1].Value, 2)
		},
	)
}

func TestDebug(t *testing.T) {
	t.Parallel()
	prove := proof.New(t)

	const input = `{"name":"first", value:1}`
	decoder := NewDecoder(strings.NewReader(input))
	decoder.SetDebug(true)
	for decoder.More() {
		var row map[string]interface{}
		prove.Err(decoder.Decode(&row))
		prove.Equal(decoder.debugData, []string{input})
	}
}
