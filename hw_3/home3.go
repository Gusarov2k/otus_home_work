// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "qwe\4\5" => "qwe45" (*)
// "qwe\45" => "qwe44444" (*)
// "qwe\\5" => "qwe\\\\\" (*)

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func check_string(name string) string {
	linkName := &name

	fullLen := len(name)
	countDigitsInWord := 0

	for _, val := range name {
		if unicode.IsDigit(val) {
			countDigitsInWord++
		}
	}
	if fullLen == countDigitsInWord {
		return ""
	} else {
		for _, val := range name {
			if unicode.IsDigit(val) {
				valAsInt, _ := strconv.Atoi(string(val))
				first := strings.Split(name, string(val))
				lastChar := len(first[0]) - 1
				repeatedChar := strings.Repeat(string(first[0][lastChar]), valAsInt-1)
				var b strings.Builder
				for i := 1; i >= 1; i-- {
					b.WriteString(first[0])
					b.WriteString(repeatedChar)
				}
				first[0] = b.String()
				*linkName = strings.Join(first, "")
			}
		}
		return name
	}
}

func main() {
	fmt.Println(check_string("a4bc2d5e"))
	fmt.Println(check_string(`qwe\45`))
}
