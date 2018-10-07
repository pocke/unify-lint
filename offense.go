package main

type Offense struct {
	Message     string  `json:"message"`
	File        string  `json:"file"`
	Type        *string `json:"type"`
	Severity    *string `json:"severity"`
	StartLine   int     `json:"start_line"`
	StartColumn *int    `json:"start_column"`
	LastLine    *int    `json:"last_line"`
	LastColumn  *int    `json:"last_column"`
}
