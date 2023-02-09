package models

import (
	"context"
	"log"
	"time"
	"tracka/pkg/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// User table
type UserCollection struct {
	*mongo.Collection
}

var coll *UserCollection

func initUsers(db *mongo.Database) {
	collection := db.Collection(collection)
	coll = &UserCollection{collection}

	var admin User
	adminDetails := config.Get().Admin
	query := bson.D{{Key: "email", Value: adminDetails.Email}}
	err := coll.FindOne(context.Background(), query).Decode(&admin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Admin not found")
			coll.CreateUser(admin.Email, admin.Password)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Admin found")
	}
}

func UserColl() *UserCollection {
	return coll
}

const collection = "users"
const cost = 10

func (coll *UserCollection) GetUser(email string) (*User, error) {
	var user User

	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (coll *UserCollection) CreateUser(email, password string) (*User, error) {
	var user User

	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}

	user = User{Email: email, createdAt: time.Now()}
	user.MakePassword(password)

	_, err = coll.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// User
type User struct {
	Email     string    `json:"email" bson:"email"`
	createdAt time.Time `bson:"createdAt"`
	Password  string    `json:"-" bson:"password"`
}

func (user *User) MakePassword(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	user.Password = string(hash)
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil
}
