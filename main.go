package main

import (
	"log"
	"os"
	"time"

	"github.com/alecthomas/kingpin"

	"github.com/alex-held/daimler-merge/pkg/merge"
)

// nolint
var (
	version = "v0.0.5"
	commit  = "snapshot"
	date    = time.Now().Format(time.RFC3339)[11:]
)

func main() {
	app := kingpin.
		New("daimler-merge", "coding task 2 for a job interview").
		Version(version).
		Author("alex-held")

	if arg := kingpin.MustParse(app.Parse(os.Args[1:])); arg == "" {
		before := merge.Ranges{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
		after := merge.Merge(before)
		logger := log.New(os.Stdout, "", 0)
		logger.Printf("Before: %v\nAfter: %v", before, after)
	}
}
