package entrypoint_extension

import _ "embed"

//go:embed Makefile.template
var _TEMPLATE_MAKEFILE []byte

//go:embed main.go.template
var _TEMPLATE_MAIN_GO []byte

type templates struct {
	Golang []byte
}

func makefileTemplate() templates {
	return templates{Golang: _TEMPLATE_MAKEFILE}
}

func mainFileTemplate() templates {
	return templates{Golang: _TEMPLATE_MAIN_GO}
}
