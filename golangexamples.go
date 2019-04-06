package golangexamples

import (
	"fmt"
	"strings"

	"github.com/ehteshamz/greetings"
)

var string1 = "Hello World!"

func ConcatSlice(sliceToConcat []byte) string {
	sliced := string(sliceToConcat)
	len := len(sliced)
	string1 := []string{}
	//j := 0
	for i := 0; i < len; i++ {
		string1 = append(string1, string(sliced[i]))
		string1 = append(string1, string('-'))
	}
	//string2 := string(string1)
	return strings.Join(string1, "")
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	newtext := ""
	message := string(sliceToEncrypt)
	message = strings.ToUpper(message)
	len := len(message)
	for i := 0; i < len; i++ {
		add := (int(message[i]) + ceaserCount)
		if add > 90 {
			add = add - 26
		}
		newtext = newtext + string(add)
	}
	fmt.Println(newtext)
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
