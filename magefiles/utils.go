package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/magefile/mage/mg"
)

type cmdOptions struct {
	args   []string
	dir    string
	stream bool
}

type cmdOption func(*cmdOptions)

func withArgs(args ...string) cmdOption {
	return func(o *cmdOptions) {
		o.args = args
	}
}

func withDir(dir string) cmdOption {
	return func(o *cmdOptions) {
		o.dir = dir
	}
}

func withStream() cmdOption {
	return func(o *cmdOptions) {
		o.stream = true
	}
}

func executeCmd(command string, options ...cmdOption) (string, error) {
	opts := &cmdOptions{}
	for _, o := range options {
		o(opts)
	}

	fmt.Printf("Executing: %s %s\n", command, strings.Join(opts.args, " "))
	cmd := exec.Command(command, opts.args...)
	if opts.dir != "" {
		cmd.Dir = opts.dir
	}

	streamOutput := mg.Verbose() || opts.stream

	var b bytes.Buffer
	if streamOutput {
		cmd.Stdout = io.MultiWriter(&b, os.Stdout)
		cmd.Stderr = io.MultiWriter(&b, os.Stderr)
	} else {
		cmd.Stdout = &b
		cmd.Stderr = &b
	}
	err := cmd.Run()
	if err != nil {
		if !streamOutput {
			fmt.Println("... failed command output:")
			fmt.Println(b.String())
		}
		return "", fmt.Errorf("error executing %s: %w", command, err)
	}
	return b.String(), nil
}
