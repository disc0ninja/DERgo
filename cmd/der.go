package main

import (
	"flag"
	"log/slog"
	"os"
	"slices"
	"time"

	"github.com/disc0ninja/dergo/internal"
	"github.com/lmittmann/tint"
)

func main() {
	// Setup logging
	w := os.Stderr
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.RFC822,
		}),
	))
	slog.Info("Welcome to \"dergo\" little program to help check that DNS is resolving as you expect")

	// Required Args:
	// File to read from. Input file should be formatted with a list (records) of dictionaries
	// as described below:
	// - name: pkg.go.dev // The DNS name of the record is the only required option
	//   expect: 10.0.0.0 // Optionnally set the `expect` and the program will ensure the record split_ansolves, and matches the expectation
	//   environments: // List of environments that this record applies to. If the environment option is not passed to the program at runtime, and a record has an environment set it will not fail if it does not split_ansolve, nor will it fail if an environment is passed as an argument but no environment is set on the record.
	//     - dev
	//     - prod

	// Optional Args:
	// Environment - Allows using the same file with different (per env)
	// settings.
	var filename, environment string
	flag.StringVar(&filename, "file", "", "file that contains list of records to try split_ansolving")
	flag.StringVar(&environment, "env", "", "environment that the program is currently being executed in")

	flag.Parse()

	if filename == "" {
		slog.Error("Filename is required!")
		os.Exit(1)
	}

	slog.Debug("Args parsed", "filename", filename, "environment", environment)

	r, err := internal.ReadRecordsFromFile(filename)
	if err != nil {
		slog.Error("Skill issue", "error", err)
		os.Exit(1)
	}

	slog.Debug("Records read from file", "recordsRead", r)

	for _, rec := range r.Records {
		if environment == "" || slices.Contains(rec.Environments, environment) {
			slog.Debug("rec", "Name", rec.Name, "Expect", rec.Expect, "Environmants", rec.Environments)
			ans, err := internal.PerformLookup(rec.Name)
			if err != nil {
				slog.Error("Something went wrong with lookup", "error", err)
			}
			slog.Info("Answer(s) found", "answer", ans)

			if rec.Expect == "" || slices.Contains(ans, rec.Expect) {
				slog.Info("Record resolves as expected", "record", rec.Name, "answers", ans, "expected", rec.Expect)
			} else {
				slog.Error("Record failes to resolve as expected", "record", rec.Name, "answers", ans, "expected", rec.Expect)
			}
		} else {
			slog.Warn("Environment not in match record environemnts", "environment", environment, "recordEnvironments", rec.Environments, "record", rec.Name)
		}
	}

}
