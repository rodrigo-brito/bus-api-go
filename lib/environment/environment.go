package environment

import (
	"fmt"
	"os"
)

func AbsPath(path string) string {
	domain := "github.com/rodrigo-brito/bus-api-go"
	return fmt.Sprintf("%s/src/%s/%s", os.Getenv("GOPATH"), domain, path)
}
