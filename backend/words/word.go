package words

import (
	"khanek/exercise-generator/core/database"

	"github.com/jinzhu/gorm"
)

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

// FindWordsByTag returns words by tag
func FindWordsByTag(tag string, limit uint) ([]*Word, error) {
	var words []*Word
	expectedTag := Tag{}
	db := database.DB()
	db.First(&expectedTag, "name = ?", tag)
	db.Limit(limit).Order("RANDOM()").Model(&expectedTag).Related(&words, "Words")
	return words, db.Error
}

// FindWordsByTagWithMinimumLenght returns words by tag but not shorted than X
func FindWordsByTagWithMinimumLenght(tag string, limit uint, minWordLenght int) ([]*Word, error) {
	var words []*Word
	expectedTag := Tag{}
	db := database.DB()
	db.First(&expectedTag, "name = ?", tag)
	db.Where("length(value) >= ?", minWordLenght).Limit(limit).Order("RANDOM()").Model(&expectedTag).Related(&words, "Words")
	return words, db.Error
}
