package main

import (
	"log"

	"github.com/magefile/mage/mg"
)

type Run mg.Namespace

func (Run) Gateway() error {
	_, err := executeCmd("go", withArgs("run", "gateway/main.go", "--dev"), withStream())
	log.Printf("Finished running gateway/main.go: %v", err)
	return nil
}

func (Run) Api() error {
	_, err := executeCmd("go", withArgs("run", "apiserver/main.go"), withStream())
	log.Printf("Finished running apiserver/main.go: %v", err)
	return nil
}
