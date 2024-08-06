package main

import (
	_ "embed"
	"errors"
	"log/slog"
	"os"
	entrypoint_extension "taku_ting/structure/internal/structure/extension/entrypoint"
	modulename_extension "taku_ting/structure/internal/structure/extension/module_name"
	"taku_ting/structure/pkg/structure"

	"github.com/spf13/pflag"
)

func main() {
	err := run()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	cliOption, err := getCLIOption()
	if err != nil {
		return err
	}

	conf, err := prepareManifest(cliOption.Prepare)
	if err != nil {
		return err
	}

	err = createStructure(conf, structure.Option{
		BaseDir:       cliOption.DestDir,
		UsingPackages: cliOption.Uses,
	})
	if err != nil {
		return err
	}

	return nil
}

type cliOption struct {
	DestDir string
	Uses    []string
	Prepare prepareOption
}

func getCLIOption() (cliOption, error) {
	dest := pflag.StringP("dest", "d", ".", "Destination for creating structure")
	module := pflag.StringP("mod", "m", "", "Module name")
	keep := pflag.Bool("keep", false, "Create .gitkeep file")
	entryPoints := pflag.StringSlice("entry", []string{"main.go"}, "Entry points to create in cmd")
	uses := pflag.StringSliceP("use", "p", []string{""}, "Packages to create")
	createDesignDoc := pflag.Bool("designdoc", false, "Create doc/DesignDoc.md")
	pflag.Parse()

	// Check cli option
	if *module == "" {
		return cliOption{}, errors.New("cli option `mod` is required and cannot be empty")
	}

	// Aggretage option
	if *createDesignDoc {
		*uses = append(*uses, "designdoc")
	}

	return cliOption{
		DestDir: *dest,
		Uses:    *uses,
		Prepare: prepareOption{
			ModuleName:    *module,
			CreateGitKeep: *keep,
			EntryPoints:   *entryPoints,
		},
	}, nil
}

type prepareOption struct {
	ModuleName    string
	CreateGitKeep bool //TODO: support
	EntryPoints   []string
}

func prepareManifest(option prepareOption) (structure.Manifest, error) {
	manifest, err := structure.ManifestCleanArchitecture()
	if err != nil {
		return manifest, err
	}

	manifest = modulename_extension.Effect(manifest, option.ModuleName)
	manifest, err = entrypoint_extension.Effect(manifest, option.EntryPoints)
	if err != nil {
		return manifest, err
	}

	return manifest, nil
}

func createStructure(manifest structure.Manifest, option structure.Option) error {
	if err := structure.Create(manifest.Structure, option); err != nil {
		return err
	}
	if err := structure.WriteFiles(manifest.Files, option); err != nil {
		return err
	}
	return nil
}
