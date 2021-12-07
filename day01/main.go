package main

import (
	"fmt"
	"io"
	"os"
)

func getSingleMeasurements(numberList []int) (result int) {
	result = 0
	for i := 1; i < len(numberList); i++ {
		if numberList[i] > numberList[i-1] {
			result++
			fmt.Println(numberList[i], " is greater than ", numberList[i-1])
			fmt.Println("Result is now: ", result)
		}
		fmt.Println("i is now: ", i)
	}

	return (result)
}

func getThreeMeasurements(numberList []int) (result int) {
	result = 0
	for i := 3; i < len(numberList); i++ {
		firstSum := (numberList[i-3] + numberList[i-2] + numberList[i-1])
		fmt.Println("The value of the first sum is", numberList[i-3], "+ ", numberList[i-2], "+ ", numberList[i-1], "= ", firstSum)
		secondSum := (numberList[i-2] + numberList[i-1] + numberList[i])
		fmt.Println("The value of the second sum is", numberList[i-2], "+ ", numberList[i-1], "+ ", numberList[i], "= ", secondSum)
		if secondSum > firstSum {
			result++
		}
		fmt.Println("result is now: ", result)
		fmt.Println("i is now: ", i)
	}

	return (result)
}
func main() {
	file, err := os.Open("inputs.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {

		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}

	//fmt.Println(getSingleMeasurements(nums))
	fmt.Println(getThreeMeasurements(nums))
}
