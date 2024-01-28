package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

type Buf mg.Namespace

func (Buf) Generate() error {
	_, err := executeCmd("rm", withArgs("-rf", "gen"))
	if err != nil {
		return fmt.Errorf("failed to run rm -rf gen: %w", err)
	}

	_, err = executeCmd("buf", withArgs("generate", "--include-imports"), withDir("protobufs"))
	if err != nil {
		return fmt.Errorf("failed to run buf generate: %w", err)
	}

	_, err = executeCmd("yarn", withArgs("generate:swagger"))
	if err != nil {
		return fmt.Errorf("failed to run yarn generate:swagger: %w", err)
	}

	return goGazelle()
}
