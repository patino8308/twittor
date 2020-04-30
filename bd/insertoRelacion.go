package bd

import (
	"context"
	"time"

	"github.com/patino8308/twittor/models"
)

/*InsertoRelacion inserta la ralacion*/
func InsertoRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(("twittores"))
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
