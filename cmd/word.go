package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zhengzhou1992/cmdtools/internal/word"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var str string
var mode int8
var desc = strings.Join([]string{
	"supported mode: ",
	"1: upper",
	"2: lower",
	"3: underscore to upper camel",
	"4: underscore to lower camel",
	"5: camel case to underscore",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word format",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderScore(str)
		default:
			log.Fatal("does't support yet, type help word for help")
		}
		log.Printf("output: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "input word")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "input mode")
}
