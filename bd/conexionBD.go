package bd

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
/*MongoCN es el objeto de conexion a la bd */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:sZ5LnEkKTrM2fMiM@cluster0-185wm.mongodb.net/test?retryWrites=true&w=majority")


/*ConectarBD es la funcion que me permite conectar la bd*/
func ConectarBD() *mongo.Client  {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa a la BD")
	return client
}
/*ChequeoConnection funcion para validar ping*/
func ChequeoConnection() int  {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil{		
		return 0
	}
	return 1
}