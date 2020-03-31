package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/signintech/gopdf"
)

type shuffledExercise struct {
	Tag string
}

func (e shuffledExercise) wordHandler(s string) string {
	return shuffle(s)
}

func (e shuffledExercise) getTitle() []string {
	return []string{fmt.Sprintf("Z rozsypanych liter proszę ułożyć wyrazy pasujące do kategorii: %s", e.Tag)}
}

func (e shuffledExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 40)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e shuffledExercise) ToPDF() (*gopdf.GoPdf, error) {
	return generateWordExercisePDF(e)
}

// NewShuffledExercise creates new exercise with masked words by tag
func NewShuffledExercise(tag string) pdf.Writer {
	return shuffledExercise{Tag: tag}
}
