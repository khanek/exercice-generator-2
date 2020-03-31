package exercises

import (
	"fmt"
	"khanek/exercise-generator/core/math"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"
	"strings"

	"github.com/signintech/gopdf"
)

const numberOfMappignLetters = 6

type letterCipherWordsExercise struct {
	Tag           string
	letterMapping map[rune]rune
}

func (e letterCipherWordsExercise) getTitle() []string {
	return []string{
		"Proszę zamienić litery w sylabach według klucza poniżej i odgadnąć zaszyfrowane wyrazy.",
		e.letterMappingToString(),
	}
}

func (e letterCipherWordsExercise) wordHandler(s string) string {
	return replaceLetters(strings.ToLower(s), e.letterMapping)
}

func (e letterCipherWordsExercise) getWords() ([]*words.Word, error) {
	words, err := words.FindWordsByTag(e.Tag, 36)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (e *letterCipherWordsExercise) generateLetterMapping() {
	choices := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'r', 's', 't', 'u', 'w', 'y', 'z'}
	lmapping := make(map[rune]rune)
	for i := 0; i < numberOfMappignLetters; i++ {
		letter1IDX := math.RandomInt(0, len(choices)-1)
		letter1 := choices[letter1IDX]
		choices = removeElement(choices, letter1IDX)
		letter2IDX := math.RandomInt(0, len(choices)-1)
		letter2 := choices[letter2IDX]
		choices = removeElement(choices, letter2IDX)
		lmapping[letter1] = letter2
		lmapping[letter2] = letter1
	}
	e.letterMapping = lmapping
}

func (e letterCipherWordsExercise) letterMappingToString() string {
	uniqueKeyValueMap := make(map[rune]rune)
	var letterPairs []string
	for k, v := range e.letterMapping {
		_, keyExists := uniqueKeyValueMap[k]
		_, valueExists := uniqueKeyValueMap[v]
		if keyExists || valueExists {
			continue
		}
		pair := fmt.Sprintf("%s%s", strings.ToUpper(string(k)), strings.ToUpper(string(v)))
		letterPairs = append(letterPairs, pair)
		uniqueKeyValueMap[k] = v
		uniqueKeyValueMap[v] = k
	}
	return strings.Join(letterPairs, " ")
}

func (e letterCipherWordsExercise) ToPDF() (*gopdf.GoPdf, error) {
	return generateWordExercisePDF(e)
}

// NewLetterCipherWordsExercise creates new exercise with replaced letters
func NewLetterCipherWordsExercise(tag string) pdf.Writer {
	e := letterCipherWordsExercise{Tag: tag}
	e.generateLetterMapping()
	return e
}
