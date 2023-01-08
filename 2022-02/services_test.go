package adventofcode2022_02_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "sebx2/adventofcode/2022-02"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AOC Suite 22_2")
}

const testData string = `A Y
B X
C Z`

const YourGestureCodePaper = "B"
const MyGestureCodeScissors = "Z"
const YourGestureCodeScissors = "C"
const StrategyCodeDraw = "Y"
const StrategyCodeWin = "Z"

var _ = Describe("Test Services Part 1", func() {
	It("Does rock win against scissors?", func() {
		Expect(DidIWin(Scissors, Rock)).To(Equal(1))
	})
	It("Do Scissors draw against scissors?", func() {
		Expect(DidIWin(Scissors, Scissors)).To(Equal(0))
	})
	It("Will gesture code scissors draw against gesture code scissors?", func() {
		var myAccount = new(ScoreAccount)
		expectedTotalScore := GestureScore[Scissors] + RoundScoreDraw
		myAccount.PlayRoundPart1(MyGestureCodeScissors, YourGestureCodeScissors)
		Expect(myAccount.TotalScore).To(Equal(expectedTotalScore))
	})
	It("Will test data sum be 8?", func() {
		r := strings.NewReader(testData)
		Expect(GetScoreOfMatch(r, 1)).To(Equal(15))
	})
})

var _ = Describe("Test Services Part 1", func() {
	It("Strategic gesture to win against rock should be paper", func() {
		Expect(GetStrategicGesture(Rock, StrategyCodeWin)).To(Equal(Paper))
	})
	It("Playing draw round with paper is supposed to result in score 5.", func() {
		var myAccount = new(ScoreAccount)
		expectedTotalScore := GestureScore[Paper] + RoundScoreDraw
		myAccount.PlayRoundPart2(YourGestureCodePaper, StrategyCodeDraw)
		Expect(myAccount.TotalScore).To(Equal(expectedTotalScore))
	})
	It("Will test data sum be 12?", func() {
		r := strings.NewReader(testData)
		Expect(GetScoreOfMatch(r, 2)).To(Equal(12))
	})
})
