package routes

import (
	"github.com/kataras/iris/v12"
	"log"
)

type iGetIndexResponse struct {
	Hello string `json:"hello"`
}

func GetIndex(context iris.Context) {
	err := context.JSON(iGetIndexResponse{Hello: "world"})
	context.StatusCode(200)
	if err != nil {
		log.Println("encountered an error while trying to send json to index route ", err)
	}
}
