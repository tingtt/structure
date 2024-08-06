package structure

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"structure/constants"

	"gopkg.in/yaml.v3"
)

type Manifest struct {
	Structure []Dir  `yaml:"structure"`
	Files     []File `yaml:"files"`
}

type Dir struct {
	Path string  `yaml:"path"`
	If   *string `yaml:"if,omitempty"`
}

type File struct {
	Path              string  `yaml:"path"`
	Content           string  `yaml:"content,omitempty"`
	EnableWithPackage *string `yaml:"if,omitempty"`
}

func ManifestCleanArchitecture() (Manifest, error) {
	var manifest Manifest

	err := yaml.Unmarshal(constants.YamlManifestCleanArchitecture(), &manifest)
	if err != nil {
		return Manifest{}, err
	}

	return manifest, nil
}

type Option struct {
	BaseDir       string
	UsingPackages []string
}

func Create(structure []Dir, opt Option) error {
	for _, s := range structure {
		if s.If != nil && !contains(opt.UsingPackages, *s.If) {
			continue
		}
		err := os.MkdirAll(path.Join(opt.BaseDir, s.Path), os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}

func WriteFiles(files []File, opt Option) error {
	for _, f := range files {
		if f.EnableWithPackage != nil && !contains(opt.UsingPackages, *f.EnableWithPackage) {
			continue
		}
		dir := filepath.Dir(f.Path)
		err := os.MkdirAll(path.Join(opt.BaseDir, dir), os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory for file: %w", err)
		}
		err = os.WriteFile(path.Join(opt.BaseDir, f.Path), []byte(f.Content), 0644)
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}
	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
