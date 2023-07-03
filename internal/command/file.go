package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type MockRunner struct {
	BasePath string
}

func (fcr MockRunner) Exec(command string, args ...string) ([]byte, error) {
	fileName := fmt.Sprintf("%s-%s.txt", command, strings.Join(args, "-"))

	if len(args) < 1 {
		return nil, fmt.Errorf("")
	}

	return os.ReadFile(filepath.Join(fcr.BasePath, fileName))
}
