package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	adventofcode2022_02 "sebx2/adventofcode/2022-02"
	"strconv"
)

func main() {
	Run202202()
}

func Run202202() {
	f, err := os.Open("./2022-02/input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	fmt.Println("Score: " + strconv.Itoa(adventofcode2022_02.GetScoreOfMatch(reader)))
}
