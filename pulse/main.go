package main

import (
	"fmt"
	"os"

	"github.com/evantbyrne/pulse"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("pulse", "Check the status of a web page.")
	url = app.Arg("url", "URL to check.").Required().String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	if err := pulse.Check(*url); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
