package exercises

import (
	"khanek/exercise-generator/core/math"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/signintech/gopdf"
)

const maxMessyWordLenght = 20

type messyExercise struct {
	Tag string
}

func (e messyExercise) getTitle() []string {
	return []string{
		"W podanych wyrazach mamy za dużo liter.",
		"Proszę wykreślić niepotrzebne litery aby powstały słowa.",
	}
}

func (e messyExercise) wordHandler(s string) string {
	wordLenght := len(s)
	if wordLenght >= 30 {
		return s
	}
	max := math.Max(maxMessyWordLenght-maxMessyWordLenght, 10)
	min := math.Min(4, max)
	return addRandomLettersToWord(s, math.RandomInt(min, max))
}

func (e messyExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTagWithMaximumLenght(e.Tag, 40, maxMessyWordLenght)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e messyExercise) ToPDF() (*gopdf.GoPdf, error) {
	return generateWordExercisePDF(e)
}

// NewMessyWordsExercise creates new exercise with masked words by tag
func NewMessyWordsExercise(tag string) pdf.Writer {
	return messyExercise{Tag: tag}
}
