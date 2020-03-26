package words

import (
	"github.com/jinzhu/gorm"
)

// FindWordsByTag returns words by tag
func FindWordsByTag(db *gorm.DB, tag string, limit uint) ([]*Word, error) {
	var words []*Word
	expectedTag := Tag{}
	db.First(&expectedTag, "name = ?", tag)
	db.Limit(limit).Order("RANDOM()").Model(&expectedTag).Related(&words, "Words")
	return words, db.Error
}
