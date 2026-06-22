// → connexion au serveur MongoDB via Mongo URI

package mongodbadapter

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitClient ouvre une connexion au serveur MongoDB et vérifie qu'il répond.
// Remplace sqliteadapter.InitDB(dbPath) : au lieu d'un fichier local,
// on se connecte à une URI (ex. mongodb://mongo:27017 dans Docker Compose).
func InitClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping confirme que le serveur est joignable (équivalent de sql.Open + première requête).
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
