package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/exercises"
	"khanek/exercise-generator/words"
)

func generateSimpleWordsExercisePDF(exerciseFactory func(words []*words.Word) pdf.Writer, wordsLimit uint) func(*gin.Context) {
	return func(c *gin.Context) {
		tag, exist := c.GetQuery("tag")
		if !exist {
			c.String(400, "Missing tag query param.")
			return
		}
		pdf, err := exercises.GenerateSimpleWordExercisePDF(tag, exerciseFactory, wordsLimit)
		if err != nil {
			c.Error(err)
			c.String(500, fmt.Sprintf("Could not generate exercise.\nError: %s", err))
			return
		}
		c.Header("Content-Disposition", "attachment; filename=simple-words-exercise.pdf")
		c.Header("Content-Type", "application/pdf")
		if err := pdf.Write(c.Writer); err != nil {
			log.Fatal(err)
		}
	}
}

func ceasar(c *gin.Context) {
	tag, exist := c.GetQuery("tag")
	if !exist {
		c.String(400, "Missing tag query param.")
		return
	}
	exerciseFactory := exercises.NewCeasarExercise
	var wordsLimit uint = 38
	_, exist = c.GetQuery("withHelp")
	if exist {
		exerciseFactory = exercises.NewCeasarWithHelpExercise
		wordsLimit = 36
	}

	pdf, err := exercises.GenerateSimpleWordExercisePDF(tag, exerciseFactory, wordsLimit)
	if err != nil {
		c.Error(err)
		c.String(500, fmt.Sprintf("Could not generate exercise.\nError: %s", err))
		return
	}
	c.Header("Content-Disposition", "attachment; filename=simple-words-exercise.pdf")
	c.Header("Content-Type", "application/pdf")
	if err := pdf.Write(c.Writer); err != nil {
		log.Fatal(err)
	}
}
