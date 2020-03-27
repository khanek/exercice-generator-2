package exercises

import (
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

type shuffledExercise struct {
	Tag string
}

func (e shuffledExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 40)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e shuffledExercise) ToPDF() (*gopdf.GoPdf, error) {
	pdfObj, err := pdf.NewPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create pdf")
	}
	var exerciseWords []*words.Word
	exerciseWords, err = e.getWords()
	if err != nil {
		return nil, err
	}
	lenWords := len(exerciseWords)
	answerWords := make([]pdf.Cell, lenWords)
	shuffledWords := make([]pdf.Cell, lenWords)
	for i, word := range exerciseWords {
		answerWords[i] = pdf.NewHalfWidthPageCell(addSpaces(word.Value))
		shuffledWords[i] = pdf.NewHalfWidthPageCell(addSpaces(shuffle(word.Value)))
	}
	// exercise page
	if err := pdf.AddCellsPage(pdfObj, shuffledWords); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	// answers page
	if err := pdf.AddCellsPage(pdfObj, answerWords); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	return pdfObj, nil
}

// NewShuffledExercise creates new exercise with masked words by tag
func NewShuffledExercise(tag string) pdf.Writer {
	return shuffledExercise{Tag: tag}
}
