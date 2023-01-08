package adventofcode2022_02

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func (s *ScoreAccount) PlayRound(yourGestureCode, myGestureCode string) {
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

func GetScoreOfMatch(reader io.Reader) (score int) {

	scanner := bufio.NewScanner(reader)

	var myAccount = new(ScoreAccount)
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		myAccount.PlayRound(round[0], round[1])
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
