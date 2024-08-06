package modulename_extension

import (
	"strings"
	"structure/internal/slice"
	"structure/pkg/structure"
)

func Effect(manifest structure.Manifest, moduleName string) structure.Manifest {
	manifest.Files = slice.Map(manifest.Files, func(file structure.File) structure.File {
		file.Content = strings.ReplaceAll(file.Content, "${MODULE_NAME}", moduleName)
		file.Content = strings.ReplaceAll(file.Content, "$(MODULE_NAME)", moduleName)
		return file
	})
	return manifest
}
