package main

import (
	"fmt"
	"io"
)

func OutputSimple(offenses []Offense, out io.Writer) error {
	for _, o := range offenses {
		msg := fmt.Sprintf("%s:%d:", o.File, o.StartLine)
		if o.StartColumn != nil {
			msg += fmt.Sprintf("%d:", *o.StartColumn)
		}
		msg += " "
		msg += o.Message
		fmt.Fprintln(out, msg)
	}
	return nil
}
