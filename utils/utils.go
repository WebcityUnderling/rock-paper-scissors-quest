package utils

import (
	"bufio"
	"errors"
	"rpsq/setup"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

type Prompter interface {
	SelectPrompt(label string, items []string) (string, error)
}

type LivePrompter struct{}

var ErrSelectInvalid = errors.New("Invalid selection.")

func (rp LivePrompter) SelectPrompt(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label:    label,
		Items:    items,
		HideHelp: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", ErrSelectInvalid
	}

	return result, nil
}

func PrintAsciiByLine(content string) {
	timeout := 75 * time.Millisecond
	if setup.Environment == "development" {
		timeout = 0
	}

	s := bufio.NewScanner(strings.NewReader(content))
	for s.Scan() {
		println(s.Text())
		time.Sleep(timeout)
	}
}
