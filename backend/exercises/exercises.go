package exercises

import (
	"khanek/exercise-generator/core/pdf"

	"github.com/pkg/errors"
	"github.com/signintech/gopdf"
)

// GenerateSimpleWordExercisePDF returns pdf object with generated exercise
func GenerateSimpleWordExercisePDF(tag string, exerciseFactory func(string) pdf.Writer) (*gopdf.GoPdf, error) {
	exercise := exerciseFactory(tag)
	pdf, err := exercise.ToPDF()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create exercise pdf")
	}
	return pdf, nil
}
