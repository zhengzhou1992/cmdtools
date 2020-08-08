package main

import (
	"log"

	"github.com/zhengzhou1992/cmdtools/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
