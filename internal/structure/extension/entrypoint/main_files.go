package entrypoint_extension

import (
	"fmt"
	"structure/internal/slice"
	"structure/pkg/structure"
)

func mainFiles(entryPoints []string) []structure.File {
	if /* entrypoint only cmd/main.go */ len(entryPoints) == 1 && entryPoints[0] == "main.go" {
		return []structure.File{{
			Path:    "cmd/main.go",
			Content: string(mainFileTemplate().Golang),
		}}
	}
	files := slice.Map(entryPoints, func(entryPointName string) structure.File {
		return structure.File{
			Path:    fmt.Sprintf("cmd/%s/main.go", entryPointName),
			Content: string(mainFileTemplate().Golang),
		}
	})
	return files
}
