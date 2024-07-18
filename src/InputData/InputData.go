package InputData

import "fmt"

type InputData struct {
}

func (inputData InputData) GetNumberOfFloors() int {
	var numberOfFloor int

	fmt.Println("How many floors does your building have ? ")
	fmt.Scanf("%d", &numberOfFloor)
	return numberOfFloor
}

func (inputData InputData) GetNumberOfUnits() int {
	var numberOfUnits int

	fmt.Println("How many units does your building have ? ")
	fmt.Scanf("%d", &numberOfUnits)
	return numberOfUnits
}

func (inputData InputData) GetFullName() string {
	var fullName string

	fmt.Println("How many floors does your building have ? ")
	fmt.Scanf("%s", &fullName)
	return fullName
}

func (inputData InputData) GetWaterCost() int {
	var getWaterCost int

	fmt.Println("How much does water cost ? ")
	fmt.Scanf("%d", &getWaterCost)
	return getWaterCost
}

func (inputData InputData) GetGasCost() int {
	var getGasCost int

	fmt.Println("How much does gas cost ? ")
	fmt.Scanf("%d", &getGasCost)
	return getGasCost
}

func (inputData InputData) GetElectricityCost() int {
	var getElectricityCost int

	fmt.Println("How much does electricity cost ? ")
	fmt.Scanf("%d", &getElectricityCost)
	return getElectricityCost
}

func (inputData InputData) GetChargeAmount() int {
	var getChargeAmount int

	fmt.Println("How much is the building charge ? ")
	fmt.Scanf("%d", &getChargeAmount)
	return getChargeAmount
}
