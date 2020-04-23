package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/signintech/gopdf"
)

type missingFirstLastLetterExercise struct {
	Tag string
}

func (e missingFirstLastLetterExercise) wordHandler(s string) string {
	runes := []rune(s)
	mask := '_'
	for i := range runes {
		if i == 0 || i == len(runes)-1 || runes[i+1] == ' ' || runes[i-1] == ' ' {
			if i == 0 || runes[i-1] != mask {
				runes[i] = mask
			}
		}
	}
	return string(runes)
}

func (e missingFirstLastLetterExercise) getTitle() []string {
	return []string{fmt.Sprintf("Proszę wpisać pierwszą i ostatnią literę w poniższych wyrazach")}
}

func (e missingFirstLastLetterExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 40)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e missingFirstLastLetterExercise) ToPDF() (*gopdf.GoPdf, error) {
	return generateWordExercisePDF(e)
}

// NewMissingFirstLastLetterExercise creates new exercise with masked words by tag
func NewMissingFirstLastLetterExercise(tag string) pdf.Writer {
	return missingFirstLastLetterExercise{Tag: tag}
}
