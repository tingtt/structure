package entrypoint_extension

import (
	"fmt"
	"structure/internal/slice"
	"structure/pkg/structure"
)

func mainFiles(entryPoints []string) []structure.File {
	files := slice.Map(entryPoints, func(entryPointName string) structure.File {
		return structure.File{
			Path:    fmt.Sprintf("cmd/%s/main.go", entryPointName),
			Content: string(mainFileTemplate().Golang),
		}
	})
	return files
}
