package dao

import (
	"TeamManagementSystem/dbConfig"
	"TeamManagementSystem/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindTeambyObjectId (objectid string) (*dto.Team, error) {
	var object dto.Team

	err := dbConfig.DATABASE.Collection("Teams").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
