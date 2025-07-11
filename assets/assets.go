package assets

import (
	"embed"
)

//go:embed *.yaml rbac/*.yaml network-policy/*.yaml
var f embed.FS

// ReadFile reads and returns the content of the named file.
func ReadFile(name string) ([]byte, error) {
	return f.ReadFile(name)
}
