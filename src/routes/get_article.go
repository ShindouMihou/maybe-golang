package routes

import (
	"Maybe/src/entities"
	"Maybe/src/modules"
	"context"
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetArticle(req iris.Context) {
	iId := req.Params().Get("id")
	collection := modules.MongoClient.Database("maybe").Collection("articles")
	id, err := primitive.ObjectIDFromHex(iId)
	if err != nil {
		req.StatusCode(iris.StatusBadRequest)
		return
	}
	var result entities.Article
	err = collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&result)
	if err != nil {
		req.StatusCode(iris.StatusNotFound)
		return
	}
	_ = req.JSON(result)
}
