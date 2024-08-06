package makefile

import (
	"fmt"
	"strings"
)

func NewTarget(target string, instructions []string) string {
	targetPrefix := strings.Join([]string{phony(target), target1stLine(target)}, "\n")
	instructionDirective := strings.Join(instructions, ("\n\t"))

	return fmt.Sprintf("%s\n\t%s", targetPrefix, instructionDirective)
}

func phony(targetName string) string {
	return fmt.Sprintf(".PHONY: %s", targetName)
}

func target1stLine(targetName string) string {
	return fmt.Sprintf("%s:", targetName)
}
