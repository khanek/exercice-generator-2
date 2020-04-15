package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

type wordWithFeaturesExercise struct {
	Tag   string
	Error error
	words []*words.Word
}

func (e *wordWithFeaturesExercise) generate() {
	e.words, e.Error = words.FindWordsByTag(e.Tag, 35)
}

func (e wordWithFeaturesExercise) getTitle() string {
	return fmt.Sprintf("Proszę dopisać 10 lub więcej cech i skojarzeń do poniższych wyrazów z kategorii: %s", e.Tag)
}

func (e wordWithFeaturesExercise) ToPDF() (*gopdf.GoPdf, error) {
	if e.Error != nil {
		return nil, errors.Wrap(e.Error, "Exercise has invalid state")
	}
	pdfObj, err := pdf.NewPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create pdf")
	}
	pdfObj.AddPage()
	const margin float64 = 20
	var positionY float64
	width := gopdf.PageSizeA4.W - (margin * 2)
	height := 21.5
	pdfObj.SetX(margin)
	pdfObj.SetY(margin)
	pdfObj.SetFont("roboto", "", 12)
	err = pdfObj.CellWithOption(&gopdf.Rect{W: width, H: height}, e.getTitle(), gopdf.CellOption{Align: gopdf.Center, Border: 0})
	if err != nil {
		return nil, errors.Wrap(err, "Error when adding title to pdf")
	}
	pdfObj.SetFont("roboto", "", 14)
	pdfObj.SetX(margin)
	positionY += height + 30
	pdfObj.SetY(positionY)
	for _, word := range e.words {
		err := pdfObj.CellWithOption(&gopdf.Rect{W: width, H: height}, fmt.Sprintf("%s:", word.Value), gopdf.CellOption{Align: gopdf.Left, Border: 0})
		if err != nil {
			return nil, errors.Wrap(err, "Error when calling pdf.CellWithOption")
		}
		positionY += height
		pdfObj.SetX(margin)
		pdfObj.SetY(positionY)
	}
	return pdfObj, nil
}

// NewCeasarWithHelpExercise creates new exercise with words
func NewWordsWithFeaturesExercise(tag string) pdf.Writer {
	e := wordWithFeaturesExercise{Tag: tag}
	e.generate()
	return e
}
