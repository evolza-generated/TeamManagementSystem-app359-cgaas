package dao

import (
	"TeamManagementSystem/dbConfig"
	"TeamManagementSystem/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "context"
    "errors"
)

func DB_FindallRoles () (*[]dto.Roles, error) {
	var objects []dto.Roles
	results, err := dbConfig.DATABASE.Collection("Roless").Find(context.Background(), bson.M{})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.Roles
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Roles")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
