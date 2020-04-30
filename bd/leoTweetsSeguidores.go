package bd

import (
	"context"
	"time"

	"github.com/patino8308/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTwetsSeguidores lee los tweets*/
func LeoTwetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittores")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":          "tweet",
			"localField":    "usuarioralacionid",
			"foreingnField": "userid",
			"as":            "tweet",
		}})

	condiciones = append(condiciones, bson.M{"$unind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true
}