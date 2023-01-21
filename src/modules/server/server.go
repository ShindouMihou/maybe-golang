package server

import (
	"Maybe/src/routes"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"log"
)

var Iris = iris.New()

func Init() error {
	Iris.Use(iris.Compression)
	Iris.Validator = validator.New()
	Iris.Get("/", routes.GetIndex)
	Iris.Get("/articles/{id}/", routes.GetArticle)
	Iris.Post("/articles", routes.PostArticle)

	err := Iris.Listen(":8880")
	if err != nil {
		return err
	}
	log.Println("running on http://localhost:8080")
	return nil
}
