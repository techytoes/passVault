package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrorMsg string
	Label    string
}

func PromptGetInput(pc PromptContent, maskingAllowed bool) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := getPromptHelper(pc, templates, validate, maskingAllowed)
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func getPromptHelper(
	pc PromptContent,
	templates *promptui.PromptTemplates,
	validate promptui.ValidateFunc,
	mask bool,
) promptui.Prompt {
	if mask {
		return promptui.Prompt{
			Label:     pc.Label,
			Templates: templates,
			Validate:  validate,
			Mask:      '*',
		}
	}

	return promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}
}
