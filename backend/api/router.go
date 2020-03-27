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
	r.GET(generateURL("/ceasar"), ceasar)
}
