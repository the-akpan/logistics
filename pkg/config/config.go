package config

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Email, Password string
}

type Static struct {
	IndexPath  string
	StaticPath string
}

type Config struct {
	Db     *mongo.Database
	PORT   string
	Admin  *User
	Static *Static
}

var config Config

func Get() *Config {
	return &config
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config.Db = connectDB()
	config.PORT = configPORT()
	config.Admin = getAdmin()
	config.Static = loadStatic()
}

func connectDB() *mongo.Database {
	log.Println("Loadin DB credentials")

	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("DB_NAME")

	log.Println("Attempting DB connection")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("DB conn failed")
		log.Fatal(err)
	}

	return client.Database(database)
}

func configPORT() string {
	log.Println("Loading app port")
	port := os.Getenv("PORT")
	if _, err := strconv.Atoi(port); err != nil {
		port = "8000"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return port
}

func getAdmin() *User {
	log.Println("Loading Admin credentials")
	user := User{
		Email:    strings.TrimSpace(os.Getenv("ADMIN_EMAIL")),
		Password: strings.TrimSpace(os.Getenv("ADMIN_PASSWORD")),
	}

	if user.Email == "" || user.Password == "" {
		log.Fatal("Admin email/password not set")
	}

	return &user
}

func loadStatic() *Static {
	return &Static{
		IndexPath:  os.Getenv("STATIC_FOLDER"),
		StaticPath: os.Getenv("INDEX_FILE"),
	}
}
