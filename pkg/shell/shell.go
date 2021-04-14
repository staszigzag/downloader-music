package shell

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
)

type Runner interface {
	Run(context.Context, string) (string, error)
}

type Shell struct{}

func NewShell() *Shell {
	return &Shell{}
}

func (s *Shell) Run(command string) (answer string, err error) {
	parts := strings.Split(command, " ")

	// The first part is the command, the rest are the args
	head := parts[0]
	args := parts[1:]
	h, err := exec.LookPath(head)
	if err != nil {
		return "", err
	}
	cmd := exec.CommandContext(context.TODO(), h, args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}
