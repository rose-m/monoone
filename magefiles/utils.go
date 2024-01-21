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

func executeCmd(command string, args ...string) (string, error) {
	fmt.Printf("Executing: %s %s\n", command, strings.Join(args, " "))
	cmd := exec.Command(command, args...)

	var b bytes.Buffer
	if mg.Verbose() {
		cmd.Stdout = io.MultiWriter(&b, os.Stdout)
		cmd.Stderr = io.MultiWriter(&b, os.Stderr)
	} else {
		cmd.Stdout = &b
		cmd.Stderr = &b
	}
	err := cmd.Run()
	if err != nil {
		if !mg.Verbose() {
			fmt.Println("... failed command output:")
			fmt.Println(b.String())
		}
		return "", fmt.Errorf("error executing %s: %w", command, err)
	}
	return b.String(), nil
}
