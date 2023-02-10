package database

import (
	"context"
	"log"
	"time"
	"tracka/pkg/config"
	"tracka/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

func initUserColl(db *mongo.Database) {
	collection := db.Collection(userCollection)
	userColl = &UserCollection{collection}

	var admin models.User
	adminDetails := config.Get().Admin
	query := bson.D{{Key: "email", Value: adminDetails.Email}}
	err := userColl.FindOne(context.Background(), query).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Admin not found")
			_, err := userColl.CreateUser(adminDetails.Email, adminDetails.Password)
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Admin found")
	}
}

// User table
type UserCollection struct {
	*mongo.Collection
}

func (coll *UserCollection) GetUser(email string) (*models.User, error) {
	var user models.User

	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (coll *UserCollection) CreateUser(email, password string) (*models.User, error) {
	var user models.User

	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}

	user = models.User{Email: email, CreatedAt: time.Now()}
	user.MakePassword(password)

	_, err = coll.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (coll *UserCollection) UpdateUser(user models.User) error {
	filter := bson.D{{Key: "email", Value: user.Email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: user.Password}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	return err
}

var userColl *UserCollection

func UserColl() *UserCollection {
	return userColl
}
