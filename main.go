package main

import (
	"fmt"
	"sqlstruct/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
