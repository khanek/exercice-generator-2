package words

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Tag is a model to group words
type Tag struct {
	gorm.Model
	Name  string  `gorm:"unique_index"`
	Words []*Word `gorm:"many2many:word_tags"`
}

func (t Tag) String() string {
	return fmt.Sprintf("<Tag: '%v'>", t.Name)
}
