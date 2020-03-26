package exercises

import (
	"khanek/exercise-generator/core/db"
	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/words"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

// GenerateSimpleWordExercisePDF returns pdf object with generated exercise
func GenerateSimpleWordExercisePDF(tag string, exerciseFactory func([]*words.Word) pdf.Writer, wordsLimit uint) (*gopdf.GoPdf, error) {
	database := db.DB()
	tx := database.Begin()
	words, err := words.FindWordsByTag(database, tag, wordsLimit)
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrap(err, "Couldn't create exercise pdf")
	}
	tx.Commit()
	if len(words) == 0 {
		return nil, errors.New("Words not found")
	}
	exercise := exerciseFactory(words)
	pdf, err := exercise.ToPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create exercise pdf")
	}
	return pdf, nil
}
