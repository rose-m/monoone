package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

type Bazel mg.Namespace

func (Bazel) Gazelle() error {
	err := protoGazelle()
	if err != nil {
		return err
	}
	err = goGazelle()
	if err != nil {
		return err
	}
	return nil
}

func protoGazelle() error {
	_, err := executeCmd("bazel", withArgs("run", "//:gazelle-update-repos-proto"))
	if err != nil {
		return fmt.Errorf("failed to run //:gazelle-update-repos: %w", err)
	}
	_, err = executeCmd("bazel", withArgs("run", "//:gazelle-proto"))
	if err != nil {
		return fmt.Errorf("failed to run //:gazelle-proto: %w", err)
	}
	return nil
}

func goGazelle() error {
	_, err := executeCmd("go", withArgs("mod", "tidy"))
	if err != nil {
		return fmt.Errorf("failed to run go mod tidy: %w", err)
	}
	_, err = executeCmd("bazel", withArgs("run", "//:gazelle-update-repos"))
	if err != nil {
		return fmt.Errorf("failed to run //:gazelle-update-repos: %w", err)
	}
	_, err = executeCmd("bazel", withArgs("run", "//:gazelle"))
	if err != nil {
		return fmt.Errorf("failed to run //:gazelle: %w", err)
	}
	return nil
}
