package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	data  []int
	xSpan int
	ySpan int
}

func NewBoard(xspan, yspan int) *Board {
	return &Board{data: make([]int, xspan*yspan), xSpan: xspan, ySpan: yspan}
}

func (b *Board) Put(x, y int, value int) {
	// optionally do some bounds checking here
	b.data[x*b.ySpan+y] = value
}
func (b *Board) Get(x, y int) int {
	// optionally do some bounds checking here
	return b.data[x*b.ySpan+y]
}
func readFileIntoLines(fileName string) (inputLines []string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
		inputLines = append(inputLines, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(inputLines)
	for i := 0; i < len(inputLines); i++ {
		fmt.Println("Input Lines [", i, "] is ", inputLines[i])
	}

	return inputLines
}

func getOrderofBalls(commaSeperatedString string) (ballSlice []string) {
	ballSlice = strings.Split(commaSeperatedString, ",")
	return ballSlice
}

func initalizeGameBoards(lines []string, boardNumber int) (initalizedBoards *Board) {
	j := (2 + (6 * boardNumber))
	GameBoard := NewBoard(5, 5)
	var strFiveLines [][]int
	for i := (2 + (6 * boardNumber)); i < j+5; i++ {
		fmt.Println("i: ", i)
		fmt.Println("lines[", i, "]: ", lines[i])
		Stringvalues := strings.Fields(lines[(i)])
		fmt.Println("Stringvalues: ", Stringvalues)
		values := make([]int, 0, len(Stringvalues))
		//fmt.Println("values: ", values)
		for _, a := range Stringvalues {
			i, _ := strconv.Atoi(a)
			values = append(values, i)
		}
		strFiveLines = append(strFiveLines, values)
		fmt.Println("values: ", values)
		fmt.Println("strFiveLines: ", strFiveLines)
	}
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			GameBoard.Put(x, y, strFiveLines[x][y])
			fmt.Println("x: ", x, " y: ", y, " strFiveLines[x][y): ", strFiveLines[x][y])
		}
	}
	fmt.Println(GameBoard)
	initalizedBoards = GameBoard

	return initalizedBoards
}
func main() {

	lines := readFileIntoLines("inputs.txt")
	//fmt.Println(lines)
	balls := getOrderofBalls(lines[0])
	fmt.Println("Balls: ", balls)
	numBoards := (len(lines) - 1) / 6
	fmt.Println("NumBoards: ", numBoards)
	mySlice := make([]*Board, numBoards)
	for i := range mySlice {
		mySlice[i] = initalizeGameBoards(lines, i)
	}
	fmt.Println(mySlice[0])
	fmt.Println(mySlice[1])
	fmt.Println(mySlice[2])
	/*
		for x := 0; x < numBoards; x++ {
			boardsSlice[x] = initalizeGameBoards(lines, x)
		}*/

}
