package exercises

import (
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

type shuffledExercise struct {
	Words []*words.Word
}

func (e shuffledExercise) ToPDF() (*gopdf.GoPdf, error) {
	pdfObj, err := pdf.NewPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create pdf")
	}
	lenWords := len(e.Words)
	words := make([]pdf.Cell, lenWords)
	reversedWords := make([]pdf.Cell, lenWords)
	for i, word := range e.Words {
		words[i] = pdf.NewHalfWidthPageCell(addSpaces(word.Value))
		reversedWords[i] = pdf.NewHalfWidthPageCell(addSpaces(shuffle(word.Value)))
	}
	// exercise page
	if err := pdf.AddCellsPage(pdfObj, reversedWords); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	// answers page
	if err := pdf.AddCellsPage(pdfObj, words); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	return pdfObj, nil
}

// NewShuffledExercise creates new exercise with masked words by tag
func NewShuffledExercise(words []*words.Word) pdf.Writer {
	return shuffledExercise{Words: words}
}
