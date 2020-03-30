package exercises

import (
	"khanek/exercise-generator/core/math"
	"log"

	"github.com/signintech/gopdf"
)

type Line struct {
	pdfptr          *gopdf.GoPdf
	text            []rune
	cursor          int
	maxWordsPerLine int
	words           []string
}

func (l Line) GetText() string {
	if l.cursor == 0 {
		return ""
	}
	if l.cursor == 1 {
		return string(l.text[0])
	}
	return string(l.text[:l.cursor-1])
}

func (l Line) GetWords() []string {
	return l.words
}

func (l Line) measureWidth() float64 {
	measureCurrentLine, err := l.pdfptr.MeasureTextWidth(l.GetText())
	if err != nil {
		log.Fatal(err)
	}
	return measureCurrentLine
}

func (l Line) pageWidth() float64 {
	return gopdf.PageSizeA4.W - (l.pdfptr.MarginLeft() + l.pdfptr.MarginRight())
}

func (l Line) runeInPageBound(r rune) bool {
	width, err := l.pdfptr.MeasureTextWidth(string(r))
	if err != nil {
		log.Fatal(err)
	}
	return width+l.measureWidth() <= l.pageWidth()
}

func (l Line) wordInPageBound(w string) bool {
	width, err := l.pdfptr.MeasureTextWidth(w)
	if err != nil {
		log.Fatal(err)
	}
	return width+l.measureWidth() <= l.pageWidth()
}

func (l *Line) addRandomLetter() bool {
	letter := randomLetter()
	if !l.runeInPageBound(letter) {
		return false
	}
	l.text[l.cursor] = letter
	l.cursor++
	return true
}

func (l Line) getRandomLetters() []rune {
	spaceLength := math.RandomInt(0, maxRandomPrefix)
	ret := make([]rune, spaceLength)
	for i := 0; i < spaceLength; i++ {
		ret[i] = randomLetter()
	}
	return ret
}

func (l *Line) addString(s string) bool {
	if l.cursor == len(l.text) {
		return false
	}

	if !l.wordInPageBound(s) {
		return false
	}

	for _, r := range []rune(s) {
		l.text[l.cursor] = r
		l.cursor++
	}
	return true
}

func (l *Line) AddRandomLetters() bool {
	w := l.getRandomLetters()
	return l.addString(string(w))
}

func (l *Line) AddWord(w string) bool {
	if len(l.words) >= l.maxWordsPerLine {
		return false
	}
	ret := l.addString(w)
	if ret {
		l.words = append(l.words, w)
	}
	return ret
}

func (l *Line) FillWithRandomLetters() {
	for l.cursor < len(l.text) {
		if !l.addRandomLetter() {
			break
		}
	}
}

func (l Line) String() string {
	return string(l.text)
}

func NewLine(p *gopdf.GoPdf, maxWordsPerLine int) Line {
	return Line{pdfptr: p, text: make([]rune, charactersPerLine), maxWordsPerLine: maxWordsPerLine, words: make([]string, 0)}
}
