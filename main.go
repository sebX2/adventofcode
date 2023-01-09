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
	if err := Run202202(1); err != nil {
		log.Fatal(err)
		return
	}
	if err := Run202202(2); err != nil {
		log.Fatal(err)
		return
	}
}

func Run202202(part int) error {
	f, err := os.Open("./2022-02/input")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(f)
	score, err := adventofcode2022_02.GetScoreOfMatch(reader, part)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	fmt.Println("Score: " + strconv.Itoa(score))
	return nil
}
