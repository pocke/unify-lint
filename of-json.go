package main

import (
	"encoding/json"
	"io"
)

func OutputJSON(offenses []Offense, out io.Writer) error {
	return json.NewEncoder(out).Encode(offenses)
}
