package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

const wordsCount = 30
const charactersPerLine = 90
const maxRandomPrefix = 90

type hiddenWordsInTextExercise struct {
	Tag string
}

func (e hiddenWordsInTextExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, wordsCount)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e hiddenWordsInTextExercise) generateLines(p *gopdf.GoPdf, w []*words.Word) []Line {
	rand.Seed(time.Now().UnixNano())
	var lines []Line
	expectedLines := 30
	wordsUsed := make(map[string]bool)
	for i := 0; i < expectedLines; i++ {
		var maxWordsPerLine int
		if i < 25 {
			maxWordsPerLine = 1
		} else {
			maxWordsPerLine = 2
		}
		line := NewLine(p, maxWordsPerLine)
		line.AddRandomLetters()
		wordsLenght := len(w)
		for j := 0; j < wordsLenght; j++ {
			var inserted bool
			word := strings.ReplaceAll(strings.ToLower(w[j].Value), " ", "")
			if _, exists := wordsUsed[word]; exists {
				continue
			}
			inserted = line.AddWord(word)
			if inserted {
				wordsUsed[word] = true
			}
			line.AddRandomLetters()
		}
		line.FillWithRandomLetters()
		lines = append(lines, line)
	}
	return lines
}

func (e hiddenWordsInTextExercise) ToPDF() (*gopdf.GoPdf, error) {
	pdfObj, err := pdf.NewPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create pdf")
	}
	var exerciseWords []*words.Word
	exerciseWords, err = e.getWords()
	if err != nil {
		return nil, err
	}

	pdfObj.AddPage()
	const margin float64 = 20
	const lineHeight float64 = 20
	if err = pdfObj.SetFont("roboto", "", 16); err != nil {
		return nil, err
	}
	pdfObj.SetMargins(margin, margin*2, margin, margin*20)
	pdfObj.SetX(margin)
	pdfObj.SetY(margin + 20)
	lines := e.generateLines(pdfObj, exerciseWords)
	answers := make([]string, 0)
	for _, line := range lines {
		for _, word := range line.GetWords() {
			answers = append(answers, word)
		}
	}
	helpText := fmt.Sprintf("W podanym ciągu liter proszę wykreślić %d wyrazów z kategorii: %v", len(answers), e.Tag)
	helpTextLines, err := pdfObj.SplitText(helpText, gopdf.PageSizeA4.W-(margin*2))
	if err != nil {
		return nil, err
	}
	for _, helpTextLine := range helpTextLines {
		pdfObj.Text(helpTextLine)
		pdfObj.SetX(margin)
		pdfObj.SetY(pdfObj.GetY() + lineHeight)
	}
	pdfObj.SetY(pdfObj.GetY() + lineHeight + 20)
	for _, line := range lines {
		pdfObj.SetX(margin)
		pdfObj.Text(line.GetText())
		pdfObj.SetY(pdfObj.GetY() + lineHeight)
	}
	pdfObj.SetX(margin)
	pdfObj.SetY(pdfObj.GetY() + 10)
	pdfObj.Text("Znalezione słowa:")
	for i := 0; i < 4; i++ {
		pdfObj.SetX(margin)
		pdfObj.SetY(pdfObj.GetY() + (lineHeight * 1.5))
		pdfObj.Line(pdfObj.GetX(), pdfObj.GetY(), gopdf.PageSizeA4.W-margin, pdfObj.GetY())
	}

	pdfObj.AddPage()
	pdfObj.SetX(margin)
	pdfObj.SetY(margin + 20)
	pdfObj.Text("Słowa ukryte w liniach:")
	pdfObj.SetX(margin)
	pdfObj.SetY(pdfObj.GetY() + margin)
	for i, line := range lines {
		lineNumber := i + 1
		words := line.GetWords()
		if len(words) > 0 {
			pdfObj.Text(fmt.Sprintf("Linia: %d, słowa: %v", lineNumber, strings.Join(line.GetWords(), ",")))
		} else {
			pdfObj.Text(fmt.Sprintf("Linia: %d, brak słów", lineNumber))
		}
		pdfObj.SetX(margin)
		pdfObj.SetY(pdfObj.GetY() + margin)
	}
	return pdfObj, nil
}

// NewHiddenWordsInTextExerciseExercise creates new exercise with masked words by tag
func NewHiddenWordsInTextExerciseExercise(tag string) pdf.Writer {
	return hiddenWordsInTextExercise{Tag: tag}
}
