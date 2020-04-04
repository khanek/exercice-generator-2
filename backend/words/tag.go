package words

import (
	"fmt"

	"github.com/jinzhu/gorm"

	qorAdmin "github.com/qor/admin"
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

func (t *Tag) AdminActions() (actions []*qorAdmin.Action) {
	return actions
}

func (t *Tag) ConfigureResource(r *qorAdmin.Resource) error {
	// r.Action(&qorAdmin.Action{
	// 	Name:  "Wygeneruj ćwiczenie (brakujące litery)",
	// 	Modes: []string{"menu_item", "edit"},
	// 	URL: func(record interface{}, context *qorAdmin.Context) string {
	// 		tag, ok := record.(*Tag)
	// 		if !ok {
	// 			return "#error"
	// 		}
	// 		queryParams := url.Values{}
	// 		queryParams.Set("tag", tag.Name)
	// 		queryParams.Set("type", strconv.Itoa(MaskedWordsExercise))
	// 		return "http://localhost:9000/api/exercises/generate?" + queryParams.Encode()
	// 	},
	// })
	// r.Action(&qorAdmin.Action{
	// 	Name:  "Wygeneruj ćwiczenie (pomieszane litery)",
	// 	Modes: []string{"menu_item", "edit"},
	// 	URL: func(record interface{}, context *qorAdmin.Context) string {
	// 		tag, ok := record.(*Tag)
	// 		if !ok {
	// 			return "#error"
	// 		}
	// 		queryParams := url.Values{}
	// 		queryParams.Set("tag", tag.Name)
	// 		queryParams.Set("type", strconv.Itoa(ShuffledWordsExercise))
	// 		return "http://localhost:9000/api/exercises/generate?" + queryParams.Encode()
	// 	},
	// })
	// r.Action(&qorAdmin.Action{
	// 	Name:  "Wygeneruj ćwiczenie (odwrócone litery)",
	// 	Modes: []string{"menu_item", "edit"},
	// 	URL: func(record interface{}, context *qorAdmin.Context) string {
	// 		tag, ok := record.(*Tag)
	// 		if !ok {
	// 			return "#error"
	// 		}
	// 		queryParams := url.Values{}
	// 		queryParams.Set("tag", tag.Name)
	// 		queryParams.Set("type", strconv.Itoa(ReversedWordsExercise))
	// 		return "http://localhost:9000/api/exercises/generate?" + queryParams.Encode()
	// 	},
	// })
	r.IndexAttrs("Name")
	return nil
}
