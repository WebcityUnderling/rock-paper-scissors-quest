package utils

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

var ErrSelectInvalid = errors.New("invalid selection")

func SelectPrompt(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label:    label,
		Items:    items,
		HideHelp: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", ErrSelectInvalid
	}

	return result, nil
}

func PrintAsciiByLine(content string) {

	s := bufio.NewScanner(strings.NewReader(content))
	for s.Scan() {
		println(s.Text())
		time.Sleep(75 * time.Millisecond)
	}
}
