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
	AnswerA() (string, error)
	AnswerB() (string, error)
}

func main() {
	var problem = flag.String("problem", "example", "input and output filename, defaults to example")
	flag.Parse()
	orchestrator := &one.Comparer{}
	err := general.Load(orchestrator, inputDirectory+"/"+*problem)
	if err != nil {
		fmt.Println(err)
		return
	}

	a, err := orchestrator.AnswerA()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = general.Save(outputDirectory+"/A/"+*problem, a)
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := orchestrator.AnswerB()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = general.Save(outputDirectory+"/B/"+*problem, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Problem " + *problem + " success!")
}
