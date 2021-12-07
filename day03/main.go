package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func findOxygenRating(nums []string) (oxygenRating string) {
	for pos := 0; pos < len(nums[0]); pos++ {
		zeroes := 0
		ones := 0
		var zeroslice []string
		var oneslice []string
		for i := 0; i < len(nums); i++ {
			fmt.Println("i: ", i, "pos :", pos)
			if nums[i][pos:pos+1] == "0" {
				zeroes++
				zeroslice = append(zeroslice, nums[i])
				fmt.Println("Zeroslice is ", zeroslice)
			} else {
				ones++
				oneslice = append(oneslice, nums[i])
				fmt.Println("Oneslice is ", oneslice)
			}
			fmt.Println("ones: ", ones, "zeroes:", zeroes)
		}
		if zeroes > ones {
			nums = zeroslice
			fmt.Println("Zeroes are greater for pos ", pos, "nums is now equal to: ", nums)
		}
		if ones >= zeroes {
			nums = oneslice
			fmt.Println("Ones are greater for pos ", pos, "nums is now equal to: ", nums)
		}
		fmt.Println("For pos: ", pos)
		if len(nums) == 1 {
			break
		}

	}
	oxygenRating = nums[0]
	return oxygenRating
}

func findScrubberRating(nums []string) (scrubberRating string) {
	for pos := 0; pos < len(nums[0]); pos++ {
		zeroes := 0
		ones := 0
		var zeroslice []string
		var oneslice []string
		for i := 0; i < len(nums); i++ {
			fmt.Println("i: ", i, "pos :", pos)
			if nums[i][pos:pos+1] == "0" {
				zeroes++
				zeroslice = append(zeroslice, nums[i])
				fmt.Println("Zeroslice is ", zeroslice)
			} else {
				ones++
				oneslice = append(oneslice, nums[i])
				fmt.Println("Oneslice is ", oneslice)
			}
			fmt.Println("ones: ", ones, "zeroes:", zeroes)
		}
		if zeroes <= ones {
			nums = zeroslice
			fmt.Println("Zeroes are lesser for pos ", pos, "nums is now equal to: ", nums)
		}
		if ones < zeroes {
			nums = oneslice
			fmt.Println("Ones are lesser for pos ", pos, "nums is now equal to: ", nums)
		}
		fmt.Println("For pos: ", pos)
		if len(nums) == 1 {
			break
		}

	}
	scrubberRating = nums[0]
	return scrubberRating
}

func findPowerConsumption(nums []string) (gamma []int, epsilon []int) {
	for pos := 0; pos < len(nums[0]); pos++ {
		zeroes := 0
		ones := 0
		for i := 0; i < len(nums); i++ {
			fmt.Println("i: ", i, "pos :", pos)
			if nums[i][pos:pos+1] == "0" {
				zeroes += 1
			} else {
				ones += 1
			}
			fmt.Println("ones: ", ones, "zeroes:", zeroes)
		}
		if zeroes > ones {
			gamma = append(gamma, 0)
			epsilon = append(epsilon, 1)
		}
		if ones > zeroes {
			gamma = append(gamma, 1)
			epsilon = append(epsilon, 0)
		}
		fmt.Println("For pos: ", pos, " gama is :", gamma, " and epsilon is: ", epsilon)

	}
	return gamma, epsilon
}

func convertToString(numberSlice []int) (numberString string) {
	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	valuesText := []string{}
	for i := range numberSlice {
		number := numberSlice[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
		fmt.Println(valuesText)
	}

	// Join our string slice.
	fmt.Println("ValuesText: ", valuesText)
	numberString = strings.Join(valuesText, "")
	fmt.Println("NumberString:", numberString)
	return numberString
}

func main() {
	file, err := os.Open("inputs.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline string
	var nums []string

	for {
		_, err := fmt.Fscanf(file, "%s\n", &perline)
		if err != nil {

			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}
	/*
			fmt.Println(nums)
			fmt.Println(len(nums[0]))
			fmt.Println(nums[0][1:2])

		gamma, epsilon := findPowerConsumption(nums)
		fmt.Println("Gamma: ", gamma)
		fmt.Println("Epsilon: ", epsilon)
		gammaString := convertToString(gamma)
		epsilonString := convertToString(epsilon)
		fmt.Println("GammaString: ", gammaString)
		fmt.Println("EpsilonString: ", epsilonString)
		g, _ := strconv.ParseInt(gammaString, 2, 64)
		e, _ := strconv.ParseInt(epsilonString, 2, 64)
		fmt.Println("g: ", g)
		fmt.Println("e: ", e)
		println("Gamma times Epsilon is: ", (e * g))
	*/
	scrubber := findScrubberRating(nums)
	s, _ := strconv.ParseInt(scrubber, 2, 64)
	oxygen := findOxygenRating(nums)
	o, _ := strconv.ParseInt(oxygen, 2, 64)
	fmt.Println("s: ", s, "o: ", o, "s times o: ", (s * o))
}
