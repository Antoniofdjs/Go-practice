/*
This file serves as practice for more GO
Topics touched here:
- slices /arrays
-bufio and Scan() Text()
- Getting more acostumed to functions in Go

Use arays to find basics stats.
	Missing features:
	- negative symbol still has a bug
	- ability to do math operations between the arrays
	- more testing
	- remove switch case for a struct( i was testing switch cases first to understand Go easier)
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func getMinValue(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return math.Round(min)
}

func getMaxValue(numbers []float64) float64{
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return math.Round(max)
}

func getMean(numbers []float64)float64{
	var sum float64 = 0
	for _, num:= range numbers{
		sum += num
	}
	return math.Round(sum/float64(len(numbers)))
}

/*Pass a []string to convert to []floats64 and fill with 0 remaining spaces.
 Based on a reference size it will append that quantity of 0
 Returns []float64 or err
 Example: [1,2] size:2 sizereference:
 */
func zeroFill(args []string, size int, sizeReference int) ([]float64, error){
	numbers, err:= parseToFloat(args)
	if err!= nil{
		return nil, err
	}
	for i := 0; i < sizeReference - size; i++{
		numbers= append(numbers, 0)
	}
	return numbers, nil
}

/*Pass a []string to convert to []floats64 and fill with the mean of its own values present.
 Based on a reference size it will append that quantity of numbers
 Returns []float64 or err
 Example: [1,2] size:2 sizereference:
 */
func meanFill(args []string, size int, sizeReference int) ([]float64, error){
	numbers, err:= parseToFloat(args)
	if err!= nil{
		return nil, err
	}
	mean:= getMean(numbers)
	for i := 0; i < sizeReference - size; i++{
		numbers= append(numbers, mean)
	}
	return numbers, nil
}


func loopBMessage(argsA []string, sizeA int, argsB []string, sizeB int) string{
	fmt.Println("\n\nArray B must be same size as A")
	fmt.Printf("Array A: %v size: %v\n", argsA, sizeA)
	fmt.Printf("Array B: %v size %v\n\n", argsB, sizeB)
	fmt.Println("For the missing values you can fill them:\n    1)Manually(Fill again all)\n    2)Mean of your values\n    3)Zeros(0)")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()
	return choice
}

func welcome(){
	fmt.Println("WELCOME TO SIMPLE GO ARRAY ANALYSER")
	fmt.Println("\nUsage: Input number values for 2 arrays of the size of your choosing.\nThe program will give analysis data for both arrays and you will be ale to perform basic operations between the 2 arrays")
	fmt.Println("\nBoth arrays need to be the same size")
	
	fmt.Printf("\n\nTo populate your arrays use the following format usage:\n10.0,20.0,30.0\n\nArray A: ")
}

// parseToFloat parses each string in args as a float.
// It prints an error message if any string cannot be parsed.
//Returns []float64 or err
func parseToFloat(args []string) ([]float64, error){
	var numbers []float64

	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("-ERROR-: '%v' is not a valid number\n", arg)
			return nil, errors.New("Try to parse non number")
		}
		numbers = append(numbers,num )
	}
	return numbers, nil
}

func main(){
	var inputA string
	var inputB string
	var sizeA, sizeB int
	var loopB bool = true
	var numbersA, numbersB []float64 = nil, nil
	var err error

	//  Object like to scan line inputs
	scanner := bufio.NewScanner(os.Stdin)

	welcome()

	// InputA
	scanner.Scan()
	inputA = scanner.Text()
	inputA = strings.TrimRight(inputA, "\n")
	if strings.ContainsAny(inputA, " \t\n\r"){
		fmt.Println("-ERROR-: Cannot have white spaces")
		os.Exit(1)
	}
	argsA := strings.Split(inputA, ",")
	sizeA = len(argsA)
	for i:= range sizeA{
		if argsA[i] ==""{
			println("-ERROR-: Empty field argument")
			os.Exit(1)
		}
	}
	// InputB
	fmt.Printf("\nArray B: ")
	scanner.Scan()
	inputB = scanner.Text()
	inputB = strings.TrimRight(inputB, "\n")
	if strings.ContainsAny(inputB, " \t\n\r"){
		fmt.Println("-ERROR- Cannot have white spaces")
		os.Exit(1)
	}
	
	argsB := strings.Split(inputB, ",")
	sizeB = len(argsB)
	for i := 0; i < sizeB; i++{
		if argsB[i] ==""{
			println("-ERROR-: Empty field argument")
			os.Exit(1)
		}
	}

	//  Display inputs
	fmt.Printf("A [%v]\nB [%v]\n", inputA, inputB)

	// Validate inputs are same size
	if sizeA != sizeB{
		if sizeA< sizeB && sizeA != sizeB{
			fmt.Println("Inside the slice if")
			fmt.Println(sizeA)
			fmt.Println(sizeB)
			argsB = argsB[:sizeA] 
			sizeB = sizeA
			numbersB , err= parseToFloat(argsB)
			if err!= nil{
				os.Exit(1)
			}
		}else{
			for loopB == true{
			choice := loopBMessage(argsA, sizeA, argsB, sizeB)

			switch choice{
			case "1":
				fmt.Printf("\nArray B: ")
				scanner.Scan()
				inputB = scanner.Text()
				if strings.ContainsAny(inputB, " \t\n\r"){
					fmt.Println("-ERROR- Cannot have white spaces")
					os.Exit(1)
				}

				argsB := strings.Split(inputB, ",")
				sizeB = len(argsB)

				// Succes, parse argsB into floats
				if sizeA == sizeB{
					loopB = false
					numbersB, err = parseToFloat(argsB)
					if err!= nil{
						os.Exit(1)
					}
				}
			case "2":
				numbersB, err = meanFill(argsB, sizeB, sizeA)
				if err!= nil{
					os.Exit(1)
				}
				sizeB = len(numbersB)
				loopB = false
			case "3":
				numbersB, err = zeroFill(argsB, sizeB, sizeA)
				if err!= nil{
					os.Exit(1)
				}
				sizeB = len(numbersB)
				loopB = false
			default:
				fmt.Println("Not valid option")
			}
		}
	}}

	// Validate arguments arguments of inputs are numbers
	numbersA, err = parseToFloat(argsA)
	if err!= nil{
		os.Exit(1)
	}
	if numbersB == nil{
	numbersB, err = parseToFloat(argsA)
	if err!= nil{
		os.Exit(1)
	}}


	println("\n\n-VALIDATED DATA-")
	fmt.Printf("Array A: %v - size: %v\n", numbersA, sizeA)
	fmt.Printf("Array B: %v - size: %v\n", numbersB, sizeB)

	// Get some statistics
	fmt.Println("\nPROCESSING DATA")
	meanA:= getMean(numbersA)
	maxA:= getMaxValue(numbersA)
	minA:= getMinValue(numbersA)
	
	meanB:= getMean(numbersB)
	maxB:= getMaxValue(numbersB)
	minB:= getMinValue(numbersB)

	// Print results
	fmt.Println("----------RESULTS----------")
	fmt.Printf("\nArray A: %v Size: %v\n",numbersA, sizeA)
	fmt.Printf("Mean A: %v Max A: %v Min A: %v\n\n", meanA, maxA, minA)

	fmt.Printf("Array B: %v Size: %v\n",numbersB, sizeB)
	fmt.Printf("Mean B: %v Max B: %v Min B: %v\n\n", meanB, maxB, minB)
}