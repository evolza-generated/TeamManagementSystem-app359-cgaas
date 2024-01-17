package dao

import (
	"TeamManagementSystem/dbConfig"
	"TeamManagementSystem/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindRolesbyObjectId (objectid string) (*dto.Roles, error) {
	var object dto.Roles

	err := dbConfig.DATABASE.Collection("Roless").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
