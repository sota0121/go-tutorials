package main

import (
	"flag"
	"fmt"

	"github.com/sota0121/go-tutorials/features/fuzzing"
	"github.com/sota0121/go-tutorials/features/restapi"
	"github.com/sota0121/go-tutorials/features/webapp"
)

type feature string

func (f feature) String() string {
	return string(f)
}

func (f *feature) Set(value string) error {
	*f = feature(value)
	return nil
}

const (
	restapiFt feature = "restapi"
	fuzzingFt feature = "fuzzing"
	webappFt  feature = "webapp"
)

func main() {
	// Command Line Arguments Parsing
	var ft feature
	ftUsage := fmt.Sprintf("feature to run (options: %s, %s, %s)", restapiFt, fuzzingFt, webappFt)
	flag.Var(&ft, "feature", ftUsage)
	flag.Parse()

	// Run the feature
	fmt.Println(">> selected feature:", ft.String())
	switch ft {
	case restapiFt:
		restapi.Main()
	case fuzzingFt:
		fuzzing.Main()
	case webappFt:
		webapp.Main()
	default:
		fmt.Println("feature is not selected")
	}
}
