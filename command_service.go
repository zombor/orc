package orc

import (
	"fmt"
	"os/exec"
)

type CommandService interface {
	Run(string, map[string]string) (string, error)
}

func NewCommandService() commandService {
	return commandService{}
}

type commandService struct{}

func (cs commandService) Run(command string, params map[string]string) (string, error) {
	fmt.Printf("Command: %#v\n", command)
	fmt.Printf("Params: %#v\n", params)

	out, err := exec.Command(command).Output()

	return string(out), err
}
