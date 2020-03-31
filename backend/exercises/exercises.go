package exercises

import (
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

type WordsExercise interface {
	wordHandler(string) string
	getTitle() []string
	getWords() ([]*words.Word, error)
}

// GenerateSimpleWordExercisePDF returns pdf object with generated exercise
func GenerateSimpleWordExercisePDF(tag string, exerciseFactory func(string) pdf.Writer) (*gopdf.GoPdf, error) {
	exercise := exerciseFactory(tag)
	pdf, err := exercise.ToPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create exercise pdf")
	}
	return pdf, nil
}

func generateWordExercisePDF(wordExercise WordsExercise) (*gopdf.GoPdf, error) {
	pdfObj, err := pdf.NewPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create pdf")
	}
	var exerciseWords []*words.Word
	exerciseWords, err = wordExercise.getWords()
	if err != nil {
		return nil, err
	}
	lenWords := len(exerciseWords)
	answerWords := make([]pdf.Cell, lenWords)
	titles := wordExercise.getTitle()
	maskedWords := make([]pdf.Cell, lenWords+len(titles))
	for i, title := range titles {
		maskedWords[i] = pdf.NewFullWidthPageCell(title)
	}
	for i, word := range exerciseWords {
		answerWords[i] = pdf.NewHalfWidthPageCell(addSpaces(word.Value))
		maskedWords[i+len(titles)] = pdf.NewHalfWidthPageCell(addSpaces(wordExercise.wordHandler((word.Value))))
	}
	// exercise page
	if err := pdf.AddCellsPage(pdfObj, maskedWords); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	// answers page
	if err := pdf.AddCellsPage(pdfObj, answerWords); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	return pdfObj, nil
}
