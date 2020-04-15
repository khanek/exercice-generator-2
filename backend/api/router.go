package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"khanek/exercise-generator/exercises"
)

const exercisesPrefixGenerate = "/api/exercises/generate"

func generateURL(path string) string {
	return fmt.Sprintf("%s%s", exercisesPrefixGenerate, path)
}

// AddUrls adds endpoint to gin Engine
func AddUrls(r *gin.Engine) {
	const exercisesPrefixGenerate = "/api/exercises/generate"
	r.GET(generateURL("/masked"), generateSimpleWordsExercisePDF(exercises.NewMaskedExercise))
	r.GET(generateURL("/reversed"), generateSimpleWordsExercisePDF(exercises.NewReversedExercise))
	r.GET(generateURL("/shuffled"), generateSimpleWordsExercisePDF(exercises.NewShuffledExercise))
	r.GET(generateURL("/hidden"), generateSimpleWordsExercisePDF(exercises.NewHiddenWordsInTextExerciseExercise))
	r.GET(generateURL("/replaced-letters"), generateSimpleWordsExercisePDF(exercises.NewLetterCipherWordsExercise))
	r.GET(generateURL("/replaced-letters-with-numbers"), generateSimpleWordsExercisePDF(exercises.NewNumberCipherWordsExercise))
	r.GET(generateURL("/messy"), generateSimpleWordsExercisePDF(exercises.NewMessyWordsExercise))
	r.GET(generateURL("/ceasar"), func(c *gin.Context) {
		exerciseFactory := exercises.NewCeasarExercise
		if _, exist := c.GetQuery("withHelp"); exist {
			exerciseFactory = exercises.NewCeasarWithHelpExercise
		}
		httpHandler := generateSimpleWordsExercisePDF(exerciseFactory)
		httpHandler(c)
	})
	r.GET(generateURL("/splitted"), generateSimpleWordsExercisePDF(exercises.NewSplittedWordsExercise))
	r.GET(generateURL("/features"), generateSimpleWordsExercisePDF(exercises.NewWordsWithFeaturesExercise))
}
