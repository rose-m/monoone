package main

import (
	"regexp"
	"strings"

	"github.com/magefile/mage/mg"
)

type Go mg.Namespace

var buildozerRegex = regexp.MustCompile(`^buildozer '([^']+)' (.*)$`)

func (Go) GenBuild() error {
	output, err := executeCmd("bazel", "run", "//:gazelle")
	if err != nil {
		return err
	}

	// check for any buildozer matches
	for _, line := range strings.Split(output, "\n") {
		matches := buildozerRegex.FindStringSubmatch(line)
		if len(matches) == 3 {
			_, err := executeCmd("buildozer", matches[1], matches[2])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
