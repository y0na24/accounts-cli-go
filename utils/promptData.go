package utils

import (
	"fmt"

	color "github.com/fatih/color"
)

func PromptData(value string) string {
	var res string

	color.Cyan(value)
	fmt.Scan(&res)

	return res
}
