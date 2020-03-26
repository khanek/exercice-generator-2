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

const wordsLimit = 40

// AddUrls adds endpoint to gin Engine
func AddUrls(r *gin.Engine) {
	const exercisesPrefixGenerate = "/api/exercises/generate"
	r.GET(generateURL("/masked"), generateSimpleWordsExercisePDF(exercises.NewMaskedExercise, wordsLimit))
	r.GET(generateURL("/reversed"), generateSimpleWordsExercisePDF(exercises.NewReversedExercise, wordsLimit))
	r.GET(generateURL("/shuffled"), generateSimpleWordsExercisePDF(exercises.NewShuffledExercise, wordsLimit))
	r.GET(generateURL("/ceasar"), ceasar)
}
