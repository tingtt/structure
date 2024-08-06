package entrypoint_extension

import (
	"bytes"
	"fmt"
	"path"
	"taku_ting/structure/internal/slice"
	"taku_ting/structure/pkg/makefile"
	"taku_ting/structure/pkg/structure"
)

func makefileForGoBuild(entryPoints []string) (structure.File, error) {
	file := bytes.Buffer{}

	if _, err := file.Write(makefileTemplate().Golang); err != nil {
		return structure.File{}, err
	}
	if _, err := file.WriteString("\n"); err != nil {
		return structure.File{}, err
	}
	if _, err := file.WriteString(makefile.NewTarget("build", buildInstructions(entryPoints))); err != nil {
		return structure.File{}, err
	}

	return structure.File{Path: "Makefile", Content: file.String()}, nil
}

func buildInstructions(entryPoints []string) []string {
	return slice.Map(entryPoints, func(entryPointName string) string {
		return buildCommand(path.Join("cmd", entryPointName, "main.go"), entryPointName)
	})
}

func buildCommand(entrypoint, dest string) string {
	return fmt.Sprintf(
		"GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o %s %s",
		dest, entrypoint,
	)
}
