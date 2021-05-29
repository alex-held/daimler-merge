package main

import (
	"os"

	"github.com/alecthomas/kingpin"
)

// nolint
var (
	version = "v0.0.2"
	commit  = "snapshot"
	date    = "snapshot"
)

func main() {
	app := kingpin.New("daimler-merge", "coding task 2 for a job interview").Version(version).Author("alex-held")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	}
}
