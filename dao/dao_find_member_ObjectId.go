package dao

import (
	"TeamManagementSystem/dbConfig"
	"TeamManagementSystem/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindMemberbyObjectId (objectid string) (*dto.Member, error) {
	var object dto.Member

	err := dbConfig.DATABASE.Collection("Members").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
