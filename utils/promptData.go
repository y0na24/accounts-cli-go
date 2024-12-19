package utils

import (
	"fmt"

	color "github.com/fatih/color"
)

func PromptData(str string) string {
	var res string

	color.Cyan(str)
	fmt.Scan(&res)

	return res
}
