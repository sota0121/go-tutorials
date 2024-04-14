package main

import (
	"flag"
	"fmt"
	"github.com/sota0121/go-tutorials/features/nw"

	myctx "github.com/sota0121/go-tutorials/features/context"
	"github.com/sota0121/go-tutorials/features/fuzzing"
	"github.com/sota0121/go-tutorials/features/ghpra"
	"github.com/sota0121/go-tutorials/features/restapi"
	"github.com/sota0121/go-tutorials/features/webapp"
)

type feature struct {
	Name string
	Desc string
	Run  func()
}

var features = map[string]*feature{}

const (
	restapiFt string = "restapi"
	fuzzingFt string = "fuzzing"
	webappFt  string = "webapp"
	contextFt string = "context"
	ghpraFt   string = "ghpra"
	nwFt      string = "nw"
)

func init() {
	RegisterFeature(restapiFt, restapi.Main, "simple rest api")
	RegisterFeature(fuzzingFt, fuzzing.Main, "fuzzing sample")
	RegisterFeature(webappFt, webapp.Main, "simple web app")
	RegisterFeature(contextFt, myctx.Main, "sample context package usage")
	RegisterFeature(ghpraFt, ghpra.Main, "sample github pull request aggregator")
	RegisterFeature(nwFt, nw.Main, "sample network communication")
}

func RegisterFeature(name string, run func(), desc string) {
	features[name] = &feature{
		Name: name,
		Run:  run,
		Desc: desc,
	}
}

func (f *feature) String() string {
	return string(f.Name)
}

func (f *feature) Set(value string) error {
	target, ok := features[value]
	if !ok {
		return fmt.Errorf("invalid feature: %s", value)
	}
	*f = *target
	return nil
}

func ListFeatureNamesInOneLine() string {
	var names string
	for name := range features {
		names += name + ","
	}
	return names
}

func PrintFeatures(fts map[string]*feature) {
	for name, ft := range fts {
		fmt.Printf("%s: %s\n", name, ft.Desc)
	}
}

func main() {
	// Command Line Arguments Parsing
	var ft feature
	ftUsage := fmt.Sprintf("feature to run (-feature -help to show help)")
	flag.Var(&ft, "feature", ftUsage)
	flag.Bool("help", false, "show help")
	flag.Parse()

	// Show help
	if flag.Lookup("help").Value.String() == "true" {
		PrintFeatures(features)
		return
	}

	// Run the feature
	fmt.Println(">> selected feature:", ft.String())
	ft.Run()
}
