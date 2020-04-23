package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/math"
	"khanek/exercise-generator/core/pdf"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

type simpleMathExercise struct {
}

func (e simpleMathExercise) wordHandler(s string) string {
	return shuffle(s)
}

func (e simpleMathExercise) getTitle() []string {
	return []string{
		fmt.Sprintf("Proszę wykonać w pamięci zadania matematyczne"),
	}
}

func (e simpleMathExercise) generateRandomOperator() string {
	operators := []string{"+", "-"}
	return operators[math.RandomInt(0, len(operators)-1)]
}

func (e simpleMathExercise) generateMathExpression() (string, int) {
	numbersCount := 2
	mathExp := make([]string, (numbersCount*2)-1)
	answer := 0
	for i := 0; i < len(mathExp); i += 2 {
		n := math.RandomInt(100, 9999)
		mathExp[i] = strconv.Itoa(n)
		if i > 0 {
			op := e.generateRandomOperator()
			mathExp[i-1] = op
			if op == "+" {
				answer += n
			} else if op == "-" {
				answer -= n
			}
		} else {
			answer += n
		}
	}
	if answer < 0 {
		mathExp = append(mathExp, "+")
		n := math.RandomInt(-answer, -answer+1000)
		mathExp = append(mathExp, strconv.Itoa(n))
		answer += n
	} else {
		mathExp = append(mathExp, "-")
		n := math.RandomInt(answer/4, answer/2)
		mathExp = append(mathExp, strconv.Itoa(n))
		answer -= n
	}
	return strings.Join(mathExp, ""), answer

}

func (e simpleMathExercise) ToPDF() (*gopdf.GoPdf, error) {
	pdfObj, err := pdf.NewPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create pdf")
	}
	lenWords := 36
	titles := e.getTitle()
	answerWords := make([]pdf.Cell, lenWords)
	maskedWords := make([]pdf.Cell, lenWords+len(titles))
	for i, title := range titles {
		maskedWords[i] = pdf.NewFullWidthPageCellWithoutBorder(title)
	}
	for i := 0; i < lenWords; i++ {
		mathExp, answer := e.generateMathExpression()
		answerWords[i] = pdf.NewHalfWidthPageCell(addSpaces(strconv.Itoa(answer)))
		maskedWords[i+len(titles)] = pdf.NewHalfWidthPageCell(addSpaces(fmt.Sprintf("%s = ", mathExp)))
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

// NewSimpleMathExercise creates new exercise with simple mathematical expression
func NewSimpleMathExercise(tag string) pdf.Writer {
	return simpleMathExercise{}
}
