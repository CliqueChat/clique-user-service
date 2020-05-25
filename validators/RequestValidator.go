package validators

import (
	"errors"
	"github.com/CliqueChat/clique-user-service/helpers"
	"github.com/CliqueChat/clique-user-service/repositories"
	"github.com/CliqueChat/clique-user-service/structs"
	"go.mongodb.org/mongo-driver/bson"
)

func ValidateUserCreationRequest(user structs.User) error {

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.UserName == "" {
		return errors.New("email is required")
	} else {
		mongoRepo := repositories.MongoRepo
		userCollection := mongoRepo.GetCollection(helpers.User)

		filter := bson.D{{"username", user.UserName}}

		userCollection.FindOne(mongoRepo.Ctx, filter)
	}

	if user.FName == "" {
		return errors.New("user's first name is required")
	}

	return nil

}
