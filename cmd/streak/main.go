package main

import (
	"fmt"
	"os"

	"github.com/simulatedsimian/strk/streak"
)

func main() {
	repos, err := streak.GetRepos("simulatedsimian")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range repos {
		fmt.Println(repos[i].Name)
	}
}