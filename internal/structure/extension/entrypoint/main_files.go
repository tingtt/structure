package entrypoint_extension

import (
	"fmt"
	"taku_ting/structure/internal/slice"
	"taku_ting/structure/pkg/structure"
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
