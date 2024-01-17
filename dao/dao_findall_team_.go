package dao

import (
	"TeamManagementSystem/dbConfig"
	"TeamManagementSystem/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "context"
    "errors"
)

func DB_FindallTeam () (*[]dto.Team, error) {
	var objects []dto.Team
	results, err := dbConfig.DATABASE.Collection("Teams").Find(context.Background(), bson.M{})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.Team
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Team")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
