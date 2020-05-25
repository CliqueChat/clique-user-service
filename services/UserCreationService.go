package services

import (
	"github.com/CliqueChat/clique-user-service/helpers"
	"github.com/CliqueChat/clique-user-service/repositories"
	"github.com/CliqueChat/clique-user-service/structs"
	"github.com/CliqueChat/clique-user-service/validators"
)

func CreateANewUser(user structs.User) error {

	error := validators.ValidateUserCreationRequest(user)

	if error != nil {
		//TODO Handle error situation
	}

	//Saving user in mongo db
	mongoRepo := repositories.MongoRepo
	userCollection := mongoRepo.GetCollection(helpers.User)
	userCollection.InsertOne(mongoRepo.Ctx, user)

	return nil

}
