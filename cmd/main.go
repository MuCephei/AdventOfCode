package main

import (
	"flag"
	"fmt"

	"internal/general"
	"internal/one"
	"internal/two"
	"internal/three"
	"internal/four"
	"internal/five"
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
	"01.txt": func () Solver {
		return &one.Orchestrator{}
	},
	"02.txt": func () Solver {
		return &two.Orchestrator{}
	},
	"03.txt": func() Solver {
		return &three.Orchestrator{}
	},
	"04.txt": func() Solver {
		return &four.Orchestrator{}
	},
	"05.txt": func() Solver {
		return &five.Orchestrator{}
	},
}

func main() {
	var problem = flag.String("problem", "example", "input and output filename, defaults to example")
	flag.Parse()
	orchestrator := problemSolvers[*problem]()
	err := general.Load(orchestrator, inputDirectory + "/" + *problem)
	if err != nil {
		fmt.Println(err)
		return
	}

	answer, err := orchestrator.Answer()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = general.Save(outputDirectory + "/" + *problem, answer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Problem " + *problem + " success!")
}
