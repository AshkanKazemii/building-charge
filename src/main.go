package main

import (
	"building-charge/src/InputData"
	"fmt"
)

func main() {

	inputData := InputData.InputData{}

	numberOfFloors := inputData.GetNumberOfFloors()

	fmt.Println(numberOfFloors)
}
