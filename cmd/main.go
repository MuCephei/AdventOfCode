package main

import (
	"flag"
	"fmt"

	"internal/general"
	"internal/one"
)

const inputDirectory string = "assets/input"
const outputDirectory string = "assets/output"

type solver interface {
	Answer() (string, error)
}

func main() {
	var problem = flag.String("problem", "example", "input and output filename, defaults to example")
	flag.Parse()
	lines, err := general.Load(inputDirectory + "/" + *problem)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Is this elegant? No. Will I bother making a nicer solution by the end of AoC? Maybe.
	orchestrator := one.NewOrchestrator(lines)

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
