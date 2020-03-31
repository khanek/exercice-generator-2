package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/math"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"
	"math/rand"
	"time"

	"github.com/signintech/gopdf"
)

type ceasarExercise struct {
	Tag          string
	Shift        int
	withAlphabet bool
}

func (e ceasarExercise) wordHandler(s string) string {
	return ceasar(s, e.Shift)
}

func (e ceasarExercise) getTitleByShift() string {
	var text string
	if e.Shift > 0 {
		text = "Cofnij"
	} else {
		text = "Przesuń"
	}
	example := "ABCD"
	return fmt.Sprintf("%s każdą literę o %d. Przykładowo: %s => (%d) => %s", text, math.Abs(e.Shift), e.wordHandler(example), math.Abs(e.Shift), example)
}

func (e ceasarExercise) getTitle() []string {
	if e.withAlphabet {
		return []string{
			fmt.Sprintf("Alfabet: %s", addSpaces(string(polishAlphabet))),
			e.getTitleByShift(),
		}
	}
	return []string{
		e.getTitleByShift(),
	}
}

func (e ceasarExercise) getWords() ([]*words.Word, error) {
	var wordsCount uint
	if e.withAlphabet {
		wordsCount = 36
	} else {
		wordsCount = 38
	}
	words, err := words.FindWordsByTag(e.Tag, wordsCount)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e ceasarExercise) ToPDF() (*gopdf.GoPdf, error) {
	return generateWordExercisePDF(e)
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
	return ceasarExercise{Tag: tag, Shift: randomShift(), withAlphabet: false}
}

// NewCeasarWithHelpExercise creates new exercise with shifted words by tag and with help texts
func NewCeasarWithHelpExercise(tag string) pdf.Writer {
	return ceasarExercise{Tag: tag, Shift: randomShift(), withAlphabet: true}
}
