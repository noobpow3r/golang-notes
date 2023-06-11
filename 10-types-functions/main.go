package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("hello!!", Uppercase))
	fmt.Println(transformString("hello!!", Prefixer("FOO_")))
	fmt.Println(transformString("hello!!", Prefixer("BAR_")))
}