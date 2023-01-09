package adventofcode2022_02

type Gesture string

const Rock Gesture = "Rock"
const Paper Gesture = "Paper"
const Scissors Gesture = "Scissors"

var gestureAssignment = map[string]Gesture{
	"A": Rock,
	"X": Rock,
	"B": Paper,
	"Y": Paper,
	"C": Scissors,
	"Z": Scissors,
}

var GestureScore = map[Gesture]int{
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
