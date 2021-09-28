package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := make(map[string]string)
	t["TX"] = "[A-Z]"
	t["NN"] = "[0-9]"
	var s string
	keys := reflect.ValueOf(t).MapKeys()

	for _, x := range keys {
		fmt.Println(x)
		// s += strings.Join(string(x), "|")
	}
	fmt.Println(s)
}
