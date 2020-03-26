package exercises

import (
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

const maskedWordPercent = 0.5

type maskedExercise struct {
	Words []*words.Word
}

func (e maskedExercise) ToPDF() (*gopdf.GoPdf, error) {
	pdfObj, err := pdf.NewPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create pdf")
	}
	lenWords := len(e.Words)
	words := make([]pdf.Cell, lenWords)
	maskedWords := make([]pdf.Cell, lenWords)
	for i, word := range e.Words {
		words[i] = pdf.NewHalfWidthPageCell(addSpaces(word.Value))
		maskedWords[i] = pdf.NewHalfWidthPageCell(addSpaces(mask(word.Value, maskedWordPercent)))
	}
	// exercise page
	if err := pdf.AddCellsPage(pdfObj, maskedWords); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	// answers page
	if err := pdf.AddCellsPage(pdfObj, words); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	return pdfObj, nil
}

// NewMaskedExercise creates new exercise with masked words by tag
func NewMaskedExercise(words []*words.Word) pdf.Writer {
	return maskedExercise{Words: words}
}
