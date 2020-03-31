package exercises

import (
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/signintech/gopdf"
)

const maskedWordPercent = 0.5

type maskedExercise struct {
	Tag string
}

func (e maskedExercise) wordHandler(s string) string {
	return mask(s, maskedWordPercent)
}

func (e maskedExercise) getTitle() []string {
	return []string{"W poniższych przykładach proszę uzupełnić brakujące litery."}
}

func (e maskedExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 38)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e maskedExercise) ToPDF() (*gopdf.GoPdf, error) {
	return generateWordExercisePDF(e)
}

// NewMaskedExercise creates new exercise with masked words by tag
func NewMaskedExercise(tag string) pdf.Writer {
	return maskedExercise{Tag: tag}
}
