package parser

import "github.com/d9ff/adaptec-raid-exporter/internal/command"

type Scanner struct {
	command command.Runner
}

func New(cmd command.Runner) *Scanner {
	return &Scanner{command: cmd}
}
