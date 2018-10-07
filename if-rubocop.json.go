package main

import (
	"encoding/json"
)

type ifRuboCopJSON struct {
	Files []rubocopFiles `json:"files"`
}

type rubocopFiles struct {
	Path     string           `json:"path"`
	Offenses []rubocopOffense `json:"offenses"`
}

type rubocopOffense struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
	CopName  string `json:"cop_name"`
	Location struct {
		StartLine   int `json:"start_line"`
		StartColumn int `json:"start_column"`
		LastLine    int `json:"last_line"`
		LastColumn  int `json:"last_column"`
	} `json:"location"`
}

func ParseRuboCopJSON(input []byte) ([]Offense, error) {
	obj := &ifRuboCopJSON{}
	err := json.Unmarshal(input, &obj)
	if err != nil {
		return nil, err
	}
	offenses := make([]Offense, 0)

	for _, f := range obj.Files {
		path := f.Path
		for _, o := range f.Offenses {
			offenses = append(offenses, Offense{
				File:        path,
				Message:     o.Message,
				Type:        &o.CopName,
				Severity:    &o.Severity,
				StartLine:   o.Location.StartLine,
				StartColumn: &o.Location.StartColumn,
				LastLine:    &o.Location.LastLine,
				LastColumn:  &o.Location.LastColumn,
			})
		}
	}

	return offenses, nil
}
