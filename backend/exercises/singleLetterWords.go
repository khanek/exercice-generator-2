package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/pdf"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

type singleLetterWordsExercise struct {
	Tag           string
	StartFromLeft bool
}

func (e singleLetterWordsExercise) getLetters() []string {
	if e.StartFromLeft {
		return []string{
			"A",
			"B",
			"C",
			"D",
			"E",
			"F",
			"G",
			"H",
			"I",
			"J",
			"K",
			"L",
			"M",
			"N",
			"O",
			"P",
			"R",
			"S",
			"T",
			"U",
			"W",
			"Z",
		}
	}
	return []string{
		"A",
		"B",
		"C",
		"Ć",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"R",
		"S",
		"T",
		"U",
		"W",
		"Z",
	}
}

func (e singleLetterWordsExercise) ToPDF() (*gopdf.GoPdf, error) {
	pdfObj, err := pdf.NewPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create pdf")
	}
	pdfObj.AddPage()
	pdfObj.SetFont("roboto", "", 16)
	const margin float64 = 20
	var positionY float64
	width := gopdf.PageSizeA4.W - (margin * 2)
	height := 24.0
	pdfObj.SetX(margin)
	pdfObj.SetY(margin)
	titles := []string{
		"Proszę wpisać jak najwięcej wyrazów kończących się kolejnymi literami alfabetu,",
		fmt.Sprintf("zgodnych z kategorią: %s", e.Tag),
	}
	for _, title := range titles {
		err = pdfObj.CellWithOption(&gopdf.Rect{W: width, H: height}, title, gopdf.CellOption{Align: gopdf.Center, Border: 0})
		if err != nil {
			return nil, errors.Wrap(err, "Error when adding title to pdf")
		}
		pdfObj.SetX(margin)
		positionY += height * 2
		pdfObj.SetY(positionY)
	}
	if !e.StartFromLeft {
		pdfObj.SetX(width - height)
	} else {
		pdfObj.SetX(margin)
	}
	positionY += height
	pdfObj.SetY(positionY)
	for _, letter := range e.getLetters() {
		err := pdfObj.CellWithOption(&gopdf.Rect{W: height, H: height}, letter, gopdf.CellOption{Align: gopdf.Center | gopdf.Middle, Border: pdf.BorderAll})
		if err != nil {
			return nil, errors.Wrap(err, "Error when calling pdf.CellWithOption")
		}
		positionY += height * 1.2
		if !e.StartFromLeft {
			pdfObj.SetX(width - height)
		} else {
			pdfObj.SetX(margin)
		}
		pdfObj.SetY(positionY)
	}
	return pdfObj, nil
}

// NewWordsStartByLetterExercise creates new exercise with words
func NewWordsStartByLetterExercise(tag string) pdf.Writer {
	e := singleLetterWordsExercise{Tag: tag, StartFromLeft: true}
	return e
}

// NewWordsEndByLetterExercise creates new exercise with words
func NewWordsEndByLetterExercise(tag string) pdf.Writer {
	e := singleLetterWordsExercise{Tag: tag, StartFromLeft: false}
	return e
}
