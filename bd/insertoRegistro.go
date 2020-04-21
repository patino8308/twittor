package bd

import (
	"context"
	"time"
	"github.com/patino8308/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

/*InsertoRegistro inserta en bd los datos del usuario*/
func InsertoRegistro(u models.Usuario)  (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)	
	defer cancel()

	db := MongoCN.Database("twittores")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil{
		return "", false, err
	}
	
	//obtener el id del registro insertado
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

