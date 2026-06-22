package mongodbadapter

import (
	"context"

	"awesomeProject/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionTodos    = "todos"
	collectionCounters = "counters"
	counterTodoID      = "todo_id"
)

// todoDocument est la forme stockée en base.
// On garde _id en int (comme l'AUTOINCREMENT SQLite) pour ne pas changer l'API HTTP.
type todoDocument struct {
	ID           primitive.ObjectID `bson:"_id"`
	Titre        string             `bson:"titre"`
	DateCreation string             `bson:"date_creation"`
	DateEcheance string             `bson:"date_echeance"`
	Completer    bool               `bson:"completer"`
}

// Repository implémente service.TodoRepository avec MongoDB.
type Repository struct {
	todos    *mongo.Collection
	counters *mongo.Collection
}

// NewRepository sélectionne la base (MONGO_DB) et les collections utilisées.
func NewRepository(client *mongo.Client, dbName string) *Repository {
	db := client.Database(dbName)
	return &Repository{
		todos:    db.Collection(collectionTodos),
		counters: db.Collection(collectionCounters),
	}
}

func (r *Repository) Create(todo domain.TodoList) (domain.TodoList, error) {
	ctx := context.Background()

	doc := todoDocument{
		ID:           primitive.NewObjectID(),
		Titre:        todo.Titre,
		DateCreation: todo.DateCreation,
		DateEcheance: todo.DateEcheance,
		Completer:    false,
	}

	_, err := r.todos.InsertOne(ctx, doc)
	if err != nil {
		return domain.TodoList{}, err
	}

	return toDomain(doc), nil
}

func (r *Repository) DisplayByDate() ([]domain.TodoList, error) {
	ctx := context.Background()

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "dateecheance", Value: 1}})

	cursor, err := r.todos.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []domain.TodoList

	for cursor.Next(ctx) {
		var doc todoDocument

		err := cursor.Decode(&doc)
		if err != nil {
			return nil, err
		}

		todos = append(todos, toDomain(doc))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *Repository) Update(id string, todo domain.TodoList) (domain.TodoList, bool, error) {
	ctx := context.Background()

	update := bson.M{
		"$set": bson.M{
			"titre":         todo.Titre,
			"date_creation": todo.DateCreation,
			"date_echeance": todo.DateEcheance,
			"completer":     todo.Completer,
		},
	}

	result, err := r.todos.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return domain.TodoList{}, false, err
	}

	if result.MatchedCount == 0 {
		return domain.TodoList{}, false, nil
	}

	todo.ID = id
	return todo, true, nil
}

func (r *Repository) Delete(id string) (bool, error) {
	ctx := context.Background()

	result, err := r.todos.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, nil
}

// nextID simule AUTOINCREMENT SQLite via une collection "counters".
// Document { _id: "todo_id", seq: N } incrémenté atomiquement à chaque création.
func (r *Repository) nextID(ctx context.Context) (int, error) {
	var counter struct {
		Seq int `bson:"seq"`
	}

	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	err := r.counters.FindOneAndUpdate(
		ctx,
		bson.M{"_id": counterTodoID},
		bson.M{"$inc": bson.M{"seq": 1}},
		opts,
	).Decode(&counter)
	if err != nil {
		return 0, err
	}

	return counter.Seq, nil
}

func toDomain(doc todoDocument) domain.TodoList {
	return domain.TodoList{
		ID:           doc.ID.Hex(),
		Titre:        doc.Titre,
		DateCreation: doc.DateCreation,
		DateEcheance: doc.DateEcheance,
		Completer:    doc.Completer,
	}
}
