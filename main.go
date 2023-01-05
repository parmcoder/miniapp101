package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

}

func match(keyword, excludeKeywords, text string) bool {
	splited := strings.FieldsFunc(keyword, func(r rune) bool {
		return r == '+'
	})

	fmt.Println(splited)

	// TODO:
	if len(excludeKeywords) > 0 {
		match, _ := regexp.MatchString(regexp.QuoteMeta(excludeKeywords), text)
		if match {
			return false
		}
	}

	for i := 0; i < len(splited); i++ {
		match, _ := regexp.MatchString(regexp.QuoteMeta(splited[i]), text)
		if !match {
			return false
		}
	}

	return true
}
