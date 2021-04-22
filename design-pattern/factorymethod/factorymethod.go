package factorymethod

import "fmt"

type Engine interface {
	Assemble()
}

type CarEngine struct {
}

func (c CarEngine) Assemble() {
	fmt.Printf("Assembling components for car engine.")
}

type TrainEngine struct {
}

func (t TrainEngine) Assemble() {
	fmt.Printf("Assembling components for train engine.")
}

func AssemblingEngine(e Engine) {
	if e != nil {
		e.Assemble()
	}
}

func GetEngine(engineType string) Engine {
	switch engineType {
	case "car":
		return &CarEngine{}
	case "train":
		return &TrainEngine{}
	default:
		fmt.Printf("Incompatible engineType")
		return nil
	}
}

func test() {
	eng := GetEngine("car")
	AssemblingEngine(eng)
}
