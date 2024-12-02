package main

import (
	"flag"
	"fmt"

	one "github.com/mucephei/adventofcode/internal/2024/one"
	general "github.com/mucephei/adventofcode/internal/general"
)

const inputDirectory string = "assets/input"
const outputDirectory string = "assets/output"

type Solver interface {
	general.DataStore
	Answer() (string, error)
}

// Is this elegant? Better then before.
// Would I put this into production? Goodness no.
// Will I bother making a nicer solution by the end of AoC? Maybe.
var problemSolvers map[string]func() Solver = map[string]func() Solver{
	"2024/01.txt": func() Solver { return &one.Orchestrator{} },
}

func main() {
	var problem = flag.String("problem", "example", "input and output filename, defaults to example")
	flag.Parse()
	orchestrator := problemSolvers[*problem]()
	err := general.Load(orchestrator, inputDirectory+"/"+*problem)
	if err != nil {
		fmt.Println(err)
		return
	}

	answer, err := orchestrator.Answer()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = general.Save(outputDirectory+"/"+*problem, answer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Problem " + *problem + " success!")
}
