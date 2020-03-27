package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/math"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

type ceasarExercise struct {
	Tag   string
	Shift int
}

func (e ceasarExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 38)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e ceasarExercise) ToPDF() (*gopdf.GoPdf, error) {
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
	maskedWords := make([]pdf.Cell, lenWords+1)
	maskedWords[0] = pdf.NewFullWidthPageCell(getHelpText(e.Shift))
	for i, word := range exerciseWords {
		answerWords[i] = pdf.NewHalfWidthPageCell(addSpaces(word.Value))
		maskedWords[i+1] = pdf.NewHalfWidthPageCell(addSpaces(ceasar(word.Value, e.Shift)))
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

type ceasarWithAlphabet struct {
	Tag   string
	Shift int
}

func (e ceasarWithAlphabet) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 36)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e ceasarWithAlphabet) ToPDF() (*gopdf.GoPdf, error) {
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
	cryptedWords := make([]pdf.Cell, lenWords+2)
	cryptedWords[0] = pdf.NewFullWidthPageCell(getAlphabetText())
	cryptedWords[1] = pdf.NewFullWidthPageCell(getHelpText(e.Shift))
	for i, word := range exerciseWords {
		answerWords[i] = pdf.NewHalfWidthPageCell(addSpaces(strings.ToUpper(word.Value)))
		cryptedWords[i+2] = pdf.NewHalfWidthPageCell(addSpaces(ceasar(strings.ToUpper(word.Value), e.Shift)))
	}
	// exercise page
	if err := pdf.AddCellsPage(pdfObj, cryptedWords); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	// answers page
	if err := pdf.AddCellsPage(pdfObj, answerWords); err != nil {
		return nil, errors.Wrap(err, "Error on add words to page")
	}
	return pdfObj, nil
}

func getAlphabetText() string {
	return fmt.Sprintf("Alfabet: %s", addSpaces(string(polishAlphabet)))
}

func getHelpText(shift int) string {
	var text string
	if shift > 0 {
		text = "Cofnij"
	} else {
		text = "Przesuń"
	}
	example := "ABCD"
	return fmt.Sprintf("%s każdą literę o %d. Przykładowo: %s => (%d) => %s", text, math.Abs(shift), ceasar(example, shift), math.Abs(shift), example)
}

func randomShift() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(2) != 0 {
		return math.RandomInt(1, 10)
	}
	return math.RandomInt(-10, -1)
}

// NewCeasarExercise creates new exercise with shifted words by tag
func NewCeasarExercise(tag string) pdf.Writer {
	return ceasarExercise{Tag: tag, Shift: randomShift()}
}

// NewCeasarWithHelpExercise creates new exercise with shifted words by tag and with help texts
func NewCeasarWithHelpExercise(tag string) pdf.Writer {
	return ceasarWithAlphabet{Tag: tag, Shift: randomShift()}
}
