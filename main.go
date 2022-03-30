package main

import (
	"embed"
	"fmt"
)

//go:embed foo
var foo embed.FS

func main() {
	b, err := foo.ReadFile("foo/foo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
