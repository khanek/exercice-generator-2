package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"khanek/exercise-generator/core/pdf"
	"khanek/exercise-generator/exercises"
)

func generateSimpleWordsExercisePDF(exerciseFactory func(string) pdf.Writer) func(*gin.Context) {
	return func(c *gin.Context) {
		tag, exist := c.GetQuery("tag")
		if !exist {
			c.String(400, "Missing tag query param.")
			return
		}
		pdf, err := exercises.GenerateSimpleWordExercisePDF(tag, exerciseFactory)
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
	_, exist = c.GetQuery("withHelp")
	if exist {
		exerciseFactory = exercises.NewCeasarWithHelpExercise
	}

	pdf, err := exercises.GenerateSimpleWordExercisePDF(tag, exerciseFactory)
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
