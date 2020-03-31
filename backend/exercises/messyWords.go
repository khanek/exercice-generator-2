package exercises

import (
	"khanek/exercise-generator/core/math"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/signintech/gopdf"
)

type messyExercise struct {
	Tag string
}

func (e messyExercise) getTitle() []string {
	return []string{"W podanych wyrazach mamy za dużo liter. Proszę wykreślić niepotrzebne litery aby powstały słowa."}
}

func (e messyExercise) wordHandler(s string) string {
	return addRandomLettersToWord(s, math.RandomInt(4, 10))
}

func (e messyExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 40)
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
