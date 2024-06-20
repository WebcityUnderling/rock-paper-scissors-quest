package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func SelectPrompt(label string, items []string) string {
	prompt := promptui.Select{
		Label:    label,
		Items:    items,
		HideHelp: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
