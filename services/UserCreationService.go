package services

import (
	"context"
	"github.com/CliqueChat/clique-user-service/helpers"
	"github.com/CliqueChat/clique-user-service/repositories"
	"github.com/CliqueChat/clique-user-service/structs"
	"github.com/CliqueChat/clique-user-service/utils"
	"time"
)

func CreateANewUser(user structs.User) error {

	error := utils.ValidateUserCreationRequest(user)

	if error != nil {
		//TODO Handle error situation
	}

	//Saving user in mongo db
	mongoRepo := repositories.MongoRepo
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	userCollection := mongoRepo.GetCollection(helpers.User)
	userCollection.InsertOne(ctx, user)

	return nil

}
