package dao

import (
	"TeamManagementSystem/dbConfig"
	"TeamManagementSystem/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "context"
    "errors"
)

func DB_FindallMember () (*[]dto.Member, error) {
	var objects []dto.Member
	results, err := dbConfig.DATABASE.Collection("Members").Find(context.Background(), bson.M{})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.Member
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Member")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
