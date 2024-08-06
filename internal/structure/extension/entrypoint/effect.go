package entrypoint_extension

import (
	"taku_ting/structure/pkg/structure"
)

func Effect(manifest structure.Manifest, entryPoints []string) (structure.Manifest, error) {
	{
		//* Append makefile
		makefile, err := makefileForGoBuild(entryPoints)
		if err != nil {
			return structure.Manifest{}, err
		}
		manifest.Files = append(manifest.Files, makefile)
	}
	{
		//* Append main.go for each entrypoint
		manifest.Files = append(manifest.Files, mainFiles(entryPoints)...)
	}
	return manifest, nil
}
