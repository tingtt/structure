# structure

## Usage

```
structure -d /path/to/module_directory --mod example \
  --entry server --entry client
```

will make structure below.

```
$ tree /path/to/module_directory
.
├── Makefile
├── README.md
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
└── internal
    ├── domain
    ├── infrastructure
    │   ├── api
    │   └── datasource
    ├── interface-adapter
    │   └── handler
    └── usecase
```

## `--help`

```
Usage of ./structure:
      --designdoc       Create doc/DesignDoc.md
  -d, --dest string     Destination for creating structure (default ".")
      --entry strings   Entry points to create in cmd (default [main.go])
      --keep            Create .gitkeep file
  -m, --mod string      Module name
  -p, --use strings     Packages to create
```

## Build

```bash
make build
```
