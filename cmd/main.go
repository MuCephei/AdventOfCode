package main

import (
	"flag"
	"fmt"

	"internal/general"
)

const inputDirectory string = "assets/input"
const outputDirectory string = "assets/output"

func main() {
	var problem = flag.String("problem", "example", "input and output filename, defaults to example")
	flag.Parse()
	lines, err := general.Load(inputDirectory + "/" + *problem)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = general.Save(outputDirectory + "/" + *problem, lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Problem " + *problem + " success!")
}
