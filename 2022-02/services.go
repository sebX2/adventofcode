package adventofcode2022_02

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strings"
)

func (s *ScoreAccount) PlayRoundPart1(yourGestureCode, myGestureCode string) {
	yourGesture := gestureAssignment[yourGestureCode]
	myGesture := gestureAssignment[myGestureCode]
	gestureScore := GestureScore[myGesture]
	var matchScore int
	//fmt.Printf("Playing %s %s agianst %s %s.\n", myGestureCode, string(myGesture), yourGestureCode, string(yourGesture))
	switch DidIWin(yourGesture, myGesture) {
	case -1:
		matchScore = RoundScoreLoss
	case 0:
		matchScore = RoundScoreDraw
	case 1:
		matchScore = RoundScoreWin
	}
	roundScore := gestureScore + matchScore
	//fmt.Printf("GestureScore %d + MatchScore %d = %d\n", gestureScore, matchScore, roundScore)
	s.TotalScore += roundScore
	//fmt.Printf("Total Score: %d\n", s.TotalScore)
}

func (s *ScoreAccount) PlayRoundPart2(yourGestureCode, strategicOutcomeCode string) error {
	yourGesture := gestureAssignment[yourGestureCode]
	myGesture, err := GetStrategicGesture(yourGesture, strategicOutcomeCode)
	if err != nil {
		return err
	}
	var matchScore int
	switch DidIWin(yourGesture, myGesture) {
	case -1:
		matchScore = RoundScoreLoss
	case 0:
		matchScore = RoundScoreDraw
	case 1:
		matchScore = RoundScoreWin
	}
	gestureScore := GestureScore[myGesture]
	roundScore := gestureScore + matchScore
	s.TotalScore += roundScore
	return nil
}

func GetStrategicGesture(yourGesture gesture, strategicOutcomeCode string) (gesture, error) {
	switch strategicOutcomeCode {
	// loose
	case "X":
		switch yourGesture {
		case Rock:
			return Scissors, nil
		case Paper:
			return Rock, nil
		case Scissors:
			return Paper, nil
		}
	// draw
	case "Y":
		switch yourGesture {
		case Rock:
			return Rock, nil
		case Paper:
			return Paper, nil
		case Scissors:
			return Scissors, nil
		}
	// win
	case "Z":
		switch yourGesture {
		case Rock:
			return Paper, nil
		case Paper:
			return Scissors, nil
		case Scissors:
			return Rock, nil
		}
	}
	return "", errors.New("UnhandeledParameter")
}

func GetScoreOfMatch(reader io.Reader, part int) (score int) {

	scanner := bufio.NewScanner(reader)

	var myAccount = new(ScoreAccount)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		switch part {
		case 1:
			myAccount.PlayRoundPart1(round[0], round[1])
		case 2:
			myAccount.PlayRoundPart2(round[0], round[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return myAccount.TotalScore
}

func DidIWin(yourGesture, myGesture gesture) int {
	switch myGesture {
	case Rock:
		switch yourGesture {
		case Rock:
			return 0
		case Paper:
			return -1
		case Scissors:
			return 1
		}
	case Paper:
		switch yourGesture {
		case Rock:
			return 1
		case Paper:
			return 0
		case Scissors:
			return -1
		}
	case Scissors:
		switch yourGesture {
		case Rock:
			return -1
		case Paper:
			return 1
		case Scissors:
			return 0
		}
	}
	return 0
}
