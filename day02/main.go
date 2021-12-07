package main

import (
	"fmt"
	"io"
	"os"
)

func plotCourse(direction []string, degree []int) (X int, Y int) {
	X = 0
	Y = 0

	for i := 0; i < len(direction); i++ {
		if direction[i] == "forward" {
			X += degree[i]
			fmt.Println("Added Pos ", degree[i], " New Postion = ", X)
		}
		if direction[i] == "down" {
			Y += degree[i]
			fmt.Println("Added Depth ", degree[i], " New Depth = ", Y)
		}
		if direction[i] == "up" {
			Y -= degree[i]
			fmt.Println("Subtracted Depth ", degree[i], " New Depth = ", Y)
		}
		fmt.Println("i = ", i)
	}
	return X, Y
}

func findAim(direction []string, degree []int) (X int, Y int) {
	X = 0
	Y = 0
	aim := 0
	for i := 0; i < len(direction); i++ {
		if direction[i] == "forward" {
			X += degree[i]
			fmt.Println("Added Pos ", degree[i], " New Postion = ", X)
			Y += (aim * degree[i])
			fmt.Println("Added Depth ", (aim * degree[i]), "which is aim: ", aim, "pos change:", degree[i], "multiplied")
			fmt.Println("New Depth is: ", Y)
		}
		if direction[i] == "down" {
			aim += degree[i]
			fmt.Println("Aim increased by ", degree[i], " New aim = ", aim)
		}
		if direction[i] == "up" {
			aim -= degree[i]
			fmt.Println("Aim decreased by ", degree[i], " New aim = ", aim)
		}
		fmt.Println("i = ", i)
	}
	return X, Y
}

func main() {
	file, err := os.Open("inputs.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int
	var direction string
	var directions []string

	for {

		_, err := fmt.Fscanf(file, "%s %d\n", &direction, &perline) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
		directions = append(directions, direction)
	}
	//fmt.Println(nums)
	//fmt.Println(directions)

	X, Y := findAim(directions, nums)
	fmt.Println("Positon = ", X)
	fmt.Println("Depth = ", Y)
	fmt.Println("Spatial Position =", (X * Y))

}
