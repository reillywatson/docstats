package main

import (
	"fmt"
	"github.com/reillywatson/docstats"
	"os"
)

func main() {
	stats, err := docstats.StatsForDir(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(stats)
}
