package command

import (
	"os/exec"
)

type Runner interface {
	Exec(command string, args ...string) ([]byte, error)
}

type ExecRunner struct{}

func (ExecRunner) Exec(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	return cmd.Output()
}
