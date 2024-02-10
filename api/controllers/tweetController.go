package controllers

import (
	entities "api/api/entities"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type tweetController struct {
	collection *mongo.Collection
}

func NewTweetController() *tweetController{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx := context.TODO()
	err = client.Connect(ctx)
	
	if err != nil{
		panic(err)
	}
	db:=client.Database("tweets")
	collection := db.Collection("tweet")
	
	return &tweetController{collection: collection}
} 

func (t *tweetController) FindAll(ctx *gin.Context){
	cursor, err := t.collection.Find(ctx, bson.D{})
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"erro ao buscar tweet"})
		return
	}
	defer cursor.Close(ctx)
	var tweets []entities.Tweet
	if err := cursor.All(ctx, &tweets);err!= nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"erro ao buscar tweet"})
		return
	}
	ctx.JSON(http.StatusOK, tweets)
}

func (t *tweetController) Create(ctx *gin.Context){
	var tweet  entities.Tweet
	if err := ctx.BindJSON(&tweet); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar JSON"})
		return
	}
	_, err := t.collection.InsertOne(ctx, tweet)
		if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir tweet"})
		return
	}

	ctx.JSON(http.StatusOK, tweet)

}

func (t *tweetController) Delete(ctx *gin.Context){
	id := ctx.Param("id")
	result, err:= t.collection.DeleteOne(ctx, bson.M{"_id":id})
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir tweet"})
		return
	}

	if result.DeletedCount == 0 {
	ctx.JSON(http.StatusNotFound, gin.H{
		"error":"tweet nao encontrado",
	})
	return
	}

	ctx.JSON(http.StatusOK,gin.H{"msg":"tweet excluido"})
}