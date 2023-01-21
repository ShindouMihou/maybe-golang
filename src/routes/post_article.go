package routes

import (
	"Maybe/src/entities"
	"Maybe/src/modules"
	"context"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func PostArticle(req iris.Context) {
	var article entities.Article
	err := req.ReadJSON(&article)
	if err != nil {
		req.StatusCode(iris.StatusBadRequest)
		_ = req.JSON(entities.RouteError{Error: entities.InvalidPayload})
		return
	}
	collection := modules.MongoClient.Database("maybe").Collection("articles")
	result, err := collection.InsertOne(context.TODO(), article)
	if err != nil {
		req.StatusCode(iris.StatusBadGateway)
		_ = req.JSON(entities.RouteError{Error: "Failed to insert to database."})

		log.Println("error occurred while trying to insert: ", err)
		return
	}
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		_ = req.JSON(
			struct {
				Id      string           `json:"id"`
				Article entities.Article `json:"article"`
			}{
				Id:      id.Hex(),
				Article: article,
			})
	}
}
