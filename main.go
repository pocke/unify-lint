package main

import (
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

type Options struct {
	inputFormat  string
	outputFormat string
}

func ParseOptions(args []string) (*Options, []string, error) {
	opt := &Options{}
	fs := pflag.NewFlagSet(args[0], pflag.ContinueOnError)

	fs.StringVarP(&opt.inputFormat, "input-format", "i", "", "Input format")
	fs.StringVarP(&opt.outputFormat, "output-format", "o", "simple", "Output format")

	if err := fs.Parse(args[1:]); err != nil {
		return nil, nil, err
	}
	return opt, fs.Args(), nil
}

func main() {
	if err := Main(os.Args); err != nil {
		panic(err) // TODO
	}
}

func Main(args []string) error {
	opt, command, err := ParseOptions(args)
	if err != nil {
		return err
	}

	var inputFormatter func([]byte) ([]Offense, error)
	switch opt.inputFormat {
	case "rubocop.json":
		inputFormatter = ParseRuboCopJSON
	default:
		return errors.Errorf("%s is unknown input formatter.", opt.inputFormat)
	}

	var outputFormatter func([]Offense, io.Writer) error
	switch opt.outputFormat {
	case "simple":
		outputFormatter = OutputSimple
	default:
		return errors.Errorf("%s is unknown output formatter.", opt.outputFormat)
	}

	out, status, err := ExecuteLint(command)
	if err != nil {
		return err
	}
	offenses, err := inputFormatter(out)
	if err != nil {
		return err
	}

	if err := outputFormatter(offenses, os.Stdout); err != nil {
		return err
	}

	os.Exit(status)
	panic("unreachable code")
}

func ExecuteLint(command []string) ([]byte, int, error) {
	c := exec.Command(command[0], command[1:]...)
	out, err := c.Output()

	var exitStatus int
	if err != nil {
		if e2, ok := err.(*exec.ExitError); ok {
			if s, ok := e2.Sys().(syscall.WaitStatus); ok {
				exitStatus = s.ExitStatus()
			} else {
				return nil, -1, errors.New("Unimplemented for system where exec.ExitError.Sys() is not syscall.WaitStatus.")
			}
		} else {
			return nil, -1, err
		}
	} else {
		exitStatus = 0
	}

	return out, exitStatus, nil
}
