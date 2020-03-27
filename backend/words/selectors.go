package words

import (
	"khanek/exercise-generator/core/database"
)

// FindWordsByTag returns words by tag
func FindWordsByTag(tag string, limit uint) ([]*Word, error) {
	var words []*Word
	expectedTag := Tag{}
	db := database.DB()
	db.First(&expectedTag, "name = ?", tag)
	db.Limit(limit).Order("RANDOM()").Model(&expectedTag).Related(&words, "Words")
	return words, db.Error
}
