package main

import (
	"flag"
	"fmt"
)

type Args struct {
	Time int
	Urls []string
}

func newCommand() Args {
	a := Args{}

	flag.IntVar(&a.Time, "t", 15, "Timeout in seconds. 15s by default.")
	flag.IntVar(&a.Time, "timeout", 15, "Timeout in seconds. 15s by default.")

	flag.Parse()

	urls := flag.Args()
	a.Urls = urls

	fmt.Println("Команда получена")
	return a
}
