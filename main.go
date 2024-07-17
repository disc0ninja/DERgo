package main

import (
	"flag"
	"github.com/lmittmann/tint"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
	"time"
)

type Records struct {
	Records []Record `yaml:"records,flow"`
}

type Record struct {
	Name         string   `yaml:"name"`
	Expect       string   `yaml:"expect,omitempty"`
	Environments []string `yaml:"environments,flow,omitempty"`
}

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
	//   expect: 10.0.0.0 // Optionnally set the `expect` and the program will ensure the record resolves, and matches the expectation
	//   environments: // List of environments that this record applies to. If the environment option is not passed to the program at runtime, and a record has an environment set it will not fail if it does not resolve, nor will it fail if an environment is passed as an argument but no environment is set on the record.
	//     - dev
	//     - prod

	// Optional Args:
	// Environment - Allows using the same file with different (per env)
	// settings.
	var filename, environment string
	flag.StringVar(&filename, "file", "", "file that contains list of records to try resolving")
	flag.StringVar(&environment, "env", "", "environment that the program is currently being executed in")

	flag.Parse()

	slog.Debug("Args parsed", "filename", filename, "environment", environment)

	r := Records{}

	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error("Skill issue", "error", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(data), &r)
	if err != nil {
		slog.Error("Skill issue", "error", err)
		os.Exit(1)
	}

	for _, rec := range r.Records {
		slog.Info("rec", "Name", rec.Name, "Expect", rec.Expect, "Environmants", rec.Environments)
	}
}
