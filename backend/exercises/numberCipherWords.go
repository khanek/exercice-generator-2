package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/math"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"
	"strings"

	"github.com/signintech/gopdf"
)

type numberCipherWordsExercise struct {
	Tag           string
	letterMapping map[rune]rune
}

func (e numberCipherWordsExercise) wordHandler(s string) string {
	return replaceLetters(strings.ToLower(s), e.letterMapping)
}

func (e numberCipherWordsExercise) getTitle() []string {
	var letterPairs []string
	for k, v := range e.letterMapping {
		pair := fmt.Sprintf("%s => %s", strings.ToUpper(string(k)), string(v))
		letterPairs = append(letterPairs, pair)
	}
	return []string{
		"Niektóre litery zostały zastąpione cyframi. Proszę rozszyfrować wyrazy na podstawie klucza poniżej.",
		strings.Join(letterPairs, ", "),
	}
}

func (e numberCipherWordsExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 36)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e *numberCipherWordsExercise) GenerateLetterMapping() {
	choiceKeys := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'r', 's', 't', 'u', 'w', 'y', 'z'}
	choicesValues := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	lmapping := make(map[rune]rune)
	for i := 0; i < 10; i++ {
		letter1IDX := math.RandomInt(0, len(choiceKeys)-1)
		key := choiceKeys[letter1IDX]
		choiceKeys = removeElement(choiceKeys, letter1IDX)
		letter2IDX := math.RandomInt(0, len(choicesValues)-1)
		value := choicesValues[letter2IDX]
		choicesValues = removeElement(choicesValues, letter2IDX)
		lmapping[key] = value
	}
	e.letterMapping = lmapping
}

func (e numberCipherWordsExercise) ToPDF() (*gopdf.GoPdf, error) {
	return generateWordExercisePDF(e)
}

// NewNumberCipherWordsExercise creates new exercise with replaced letters
func NewNumberCipherWordsExercise(tag string) pdf.Writer {
	e := numberCipherWordsExercise{Tag: tag}
	e.GenerateLetterMapping()
	return e
}
