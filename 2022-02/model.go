package adventofcode2022_02

type gesture string

const Rock gesture = "Rock"
const Paper gesture = "Paper"
const Scissors gesture = "Scissors"

var gestureAssignment = map[string]gesture{
	"A": Rock,
	"X": Rock,
	"B": Paper,
	"Y": Paper,
	"C": Scissors,
	"Z": Scissors,
}

var GestureScore = map[gesture]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

type ScoreAccount struct {
	TotalScore int
}

const RoundScoreLoss = 0
const RoundScoreDraw = 3
const RoundScoreWin = 6
