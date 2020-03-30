package words

import "github.com/jinzhu/gorm"

// Word is a model to store words
type Word struct {
	gorm.Model
	Value string `gorm:"unique_index"`
	Tags  []*Tag `gorm:"many2many:word_tags"`
}

func (w Word) String() string {
	return w.Value
}

func WordsToStrings(words []*Word) []string {
	wordsLength := len(words)
	ret := make([]string, wordsLength)
	for i := 0; i < wordsLength; i++ {
		ret[i] = words[i].Value
	}
	return ret
}
