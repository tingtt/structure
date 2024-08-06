package constants

import _ "embed"

//go:embed clean_architecture.yaml
var yamlManifestCleanArchitecture []byte

func YamlManifestCleanArchitecture() []byte {
	return yamlManifestCleanArchitecture
}
