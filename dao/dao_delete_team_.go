package dao

import (
	"TeamManagementSystem/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteTeam (objectid string)  error {

  
        result, err := dbConfig.DATABASE.Collection("Teams").DeleteOne(context.Background(), bson.M{"objectid": objectid})
        if err != nil {
            return err
        }
        if result.DeletedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
    
}

