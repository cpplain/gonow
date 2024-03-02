package main

import (
	"fmt"
	"os"

	"github.com/cpplain/gonow/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
