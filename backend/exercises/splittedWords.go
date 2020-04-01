package exercises

import (
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

type splittedWordsExercise struct {
	Tag string
}

func (e splittedWordsExercise) wordHandler(s string) string {
	return s
}

func (e splittedWordsExercise) getTitle() []string {
	return []string{"Proszę połączyć tak, aby powstały słowa"}
}

func (e splittedWordsExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTagWithMinimumLenght(e.Tag, 19, 4)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e splittedWordsExercise) splitRunesByHalf(runes []rune) (string, string) {
	half := int(math.Floor(float64(len(runes)) / float64(2.0)))
	return string(runes[:half]), string(runes[half:])
}

func (e splittedWordsExercise) splitWord(s string) (string, string) {
	if len(s) == 4 {
		return s[:2], s[2:]
	} else if len(s) < 3 {
		return s, ""
	}
	words := strings.Split(s, " ")
	if len(words) >= 2 {
		return words[0], strings.Join(words[1:], " ")
	}
	sRunes := []rune(s)
	vowelsIndexes := make([]int, 0)
	for i := 0; i < len(sRunes); i++ {
		r := sRunes[i]
		if isVowel(r) {
			if i > 0 {
				previousVowel := sRunes[i-1]
				if r == 'e' && previousVowel == 'i' {
					vowelsIndexes[len(vowelsIndexes)-1] = i
					continue
				}
				if r == 'i' && (previousVowel == 'e' || previousVowel == 'a' || previousVowel == 'ą' || previousVowel == 'o') {
					vowelsIndexes[len(vowelsIndexes)-1] = i
					continue
				}
			}
			vowelsIndexes = append(vowelsIndexes, i)
		}
	}
	secondVowelIdx := vowelsIndexes[int(math.Floor(float64(len(vowelsIndexes)))/2.0)]
	if len(vowelsIndexes) < 2 || (secondVowelIdx+1)*2 < len(sRunes) {
		return e.splitRunesByHalf(sRunes)
	}
	return string(sRunes[:secondVowelIdx-1]), string(sRunes[secondVowelIdx-1:])
}

func (e splittedWordsExercise) ToPDF() (*gopdf.GoPdf, error) {
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
	maskedWordsFirstColumn := make([]pdf.Cell, lenWords)
	maskedWordsSecondColumn := make([]pdf.Cell, lenWords)
	for i := 0; i < len(exerciseWords); i++ {
		word := exerciseWords[i]
		answerWords[i] = pdf.NewHalfWidthPageCell(addSpaces(word.Value))
		firstPart, secondPart := e.splitWord(word.Value)
		maskedWordsFirstColumn[i] = pdf.NewHalfWidthPageCellWithoutBorder(addSpaces(firstPart))
		maskedWordsSecondColumn[i] = pdf.NewHalfWidthPageCellWithoutBorder(addSpaces(secondPart))
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(maskedWordsFirstColumn), func(i, j int) {
		maskedWordsFirstColumn[i], maskedWordsFirstColumn[j] = maskedWordsFirstColumn[j], maskedWordsFirstColumn[i]
	})
	rand.Shuffle(len(maskedWordsSecondColumn), func(i, j int) {
		maskedWordsSecondColumn[i], maskedWordsSecondColumn[j] = maskedWordsSecondColumn[j], maskedWordsSecondColumn[i]
	})
	maskedWords := make([]pdf.Cell, len(maskedWordsFirstColumn)+len(maskedWordsSecondColumn)+1)
	maskedWords[0] = pdf.NewFullWidthPageCell("Proszę połącz ze sobą lewą stronę z prawą aby postały słowa")
	for i := 0; i < len(maskedWordsFirstColumn); i++ {
		maskedWords[(i*2)+1] = maskedWordsFirstColumn[i]
		maskedWords[(i*2)+2] = maskedWordsSecondColumn[i]
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

// NewSplittedWordsExercise creates new exercise with masked words by tag
func NewSplittedWordsExercise(tag string) pdf.Writer {
	return splittedWordsExercise{Tag: tag}
}
