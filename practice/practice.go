/*
This is my intro practice for GO.
It consists in making a simple console for converting
either temperatures or weight.
The purpose of this task was to get familiarized with GO

Concepts applied:
-Conditionals
-Prints
-Error objects
-Scan
-loops
-type and map
-Switch Case
*/
package main

import (
	"fmt"
	"strconv"
)

type ConverterFunction func()(error)

func greeting()(string, error){
	var input string
	fmt.Println("\n\nChoose a converter option by typing number:\n1)Temperature\n2)Weight\nor type 'exit'")
	fmt.Scan(&input)
	fmt.Println("\nOption chosen was: ", input)
	return input, nil
}

func weightConverter()(error){
	var input string
	var inputWeight string
	var result float64
	var Error error = fmt.Errorf("No valid option selected") // Default error of this function

	fmt.Println("Choose:\n1)Kilograms(kg) -> Pounds(Lb)\n2)Pounds(Lb) -> Kilograms(kg)")
	fmt.Scan(&input)
	switch input{
	case "1":
		fmt.Printf("Kilograms: ")
		fmt.Scan(&inputWeight)
		weight, err := strconv.ParseFloat(inputWeight, 64)
		if err!= nil || weight < 0{
			if weight < 0{
				err = fmt.Errorf("Negative weight not allowed")
			}
			fmt.Println(err)
			result = 0
			Error = err
			return Error
		}
		result = weight * 2.20462
		fmt.Println("--------------------")
		fmt.Printf("# %v kg -> %v Lb\n", weight, result)
		fmt.Println("--------------------")
		Error = nil
	case "2":
		fmt.Printf("Pounds: ")
		fmt.Scan(&inputWeight)
		weight, err := strconv.ParseFloat(inputWeight, 64)
		if err!= nil || weight < 0{
			if weight < 0{
				err = fmt.Errorf("Negative weight not allowed")
			}
			fmt.Println(err)
			result = 0
			Error = err
			return Error
		}
		result = weight / 2.20462
		fmt.Println("--------------------")
		fmt.Printf("# %v Lb -> %v kg\n", weight, result)
		fmt.Println("--------------------")
		Error = nil
}
return Error
}

func tempConverter()(error){
	var input string
	var tempInput string
	var result float64
	var Error error = fmt.Errorf("No valid option was selected") // Default error of this function

	fmt.Println("Choose:\n1)Celsius -> Farenheit\n2)Farenheit -> Celsius")
	fmt.Scan(&input)

	if input == "1"{
		fmt.Printf("Celsius temperature: ")
		fmt.Scan(&tempInput)
		temp, err := strconv.ParseFloat(tempInput, 64)
		if err!= nil || temp < 0{
			fmt.Println(err)
			result = 0
			Error = err
		}
		result = (temp*(9.0/5.0)) + 32
		fmt.Println("--------------------")
		fmt.Printf("# %v C -> %v F\n", temp, result)
		fmt.Println("--------------------")
		Error = nil
	}
	if input == "2"{
		fmt.Printf("Farenheit temperature: ")
		fmt.Scan(&tempInput)
		temp, err := strconv.ParseFloat(tempInput, 64)
		if err!= nil{
			fmt.Println(err)
			result = 0
			Error = err
		}
		result = (temp - 32)*(5.0/9.0)
		fmt.Println("--------------------")
		fmt.Printf("# %v F -> %v C\n", temp, result)
		fmt.Println("--------------------")
		Error = nil
	}
	return Error
}

func main(){
	var loop = true
	converters := map[string]ConverterFunction{
		"1": tempConverter,
		"2": weightConverter,
	}

	fmt.Println("\nWELCOME")
	for loop == true{
		input, err := greeting()

		if input == "exit"{
			break
		}
		if err!=nil{
			fmt.Println("Error:", err)
		}else{
		function, exists:= converters[input]
		if !exists{
			fmt.Println("No valid option chosen, try again.")
		}else{
			err := function()
			if err!= nil{
				fmt.Println("ERROR:", err)
			}
		}
	}
	}

}
