package exercises

import (
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/signintech/gopdf"
)

type reversedExercise struct {
	Tag string
}

func (e reversedExercise) wordHandler(s string) string {
	return reverse(s)
}

func (e reversedExercise) getTitle() []string {
	return []string{"Poniższe wyrazy są zapisane wspak. O jakie wyrazy chodzi?"}
}

func (e reversedExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 38)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e reversedExercise) ToPDF() (*gopdf.GoPdf, error) {
	return generateWordExercisePDF(e)
}

// NewReversedExercise creates new exercise with masked words by tag
func NewReversedExercise(tag string) pdf.Writer {
	return reversedExercise{Tag: tag}
}
