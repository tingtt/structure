package entrypoint_extension

import (
	"bytes"
	"fmt"
	"path"
	"structure/internal/slice"
	"structure/pkg/makefile"
	"structure/pkg/structure"
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
	if /* entrypoint only cmd/main.go */ len(entryPoints) == 1 && entryPoints[0] == "main.go" {
		entryPointPath := "cmd/main.go"
		return []string{buildCommand(entryPointPath, nil /* dest */)}
	}
	return slice.Map(entryPoints, func(entryPointName string) string {
		entryPointPath := path.Join("cmd", entryPointName, "main.go")
		dest := &entryPointName
		return buildCommand(entryPointPath, dest)
	})
}

func buildCommand(entrypoint string, dest *string) string {
	if dest != nil {
		return fmt.Sprintf(
			"GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o %s %s",
			*dest, entrypoint,
		)
	}
	return fmt.Sprintf("GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build %s", entrypoint)
}
