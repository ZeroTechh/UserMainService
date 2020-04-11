package mainDB

import (
	"context"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/VelocityCore/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

// New returns a new mainDB handler struct
func New() *MainDB {
	mainDB := MainDB{}
	mainDB.init()
	return &mainDB
}

// MainDB is used to handle user mainDB data
type MainDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// init initializes client and database
func (mainDB *MainDB) init() {
	mainDB.client = utils.CreateMongoDB(dbConfig.Str("address"), log)
	mainDB.database = mainDB.client.Database(dbConfig.Str("db"))
	mainDB.collection = mainDB.database.Collection(dbConfig.Str("collection"))
}

// exists checks if user with certain field exists
func (mainDB MainDB) exists(filter types.Main) bool {
	return mainDB.Get(filter) != types.Main{}
}

// uniqueFieldsExists checks if unique fields such as username exist
func (mainDB MainDB) uniqueFieldsExists(data types.Main) string {
	usernameExists := mainDB.exists(types.Main{Username: data.Username})
	emailExists := mainDB.exists(types.Main{Email: data.Email})
	if usernameExists && data.Username != "" {
		return messages.Str("usernameExists")
	} else if emailExists && data.Email != "" {
		return messages.Str("emailExists")
	}
	return ""
}

// GenerateID is used to generate an user id
func (mainDB MainDB) GenerateID() string {
	userIDExists := true
	var userID uuid.UUID

	for userIDExists {
		userID, _ = uuid.NewRandom()
		userIDExists = mainDB.exists(types.Main{UserID: userID.String()})
	}

	return userID.String()
}

// Create is used to add new mainDB data
func (mainDB MainDB) Create(data types.Main) string {
	if !isDataValid(data) {
		return messages.Str("invalidUserData")
	}

	msg := mainDB.uniqueFieldsExists(data)
	if msg != "" {
		return msg
	}

	mainDB.collection.InsertOne(context.TODO(), data)
	return ""
}

// Get is used to a users data
func (mainDB MainDB) Get(filter types.Main) (data types.Main) {
	mainDB.collection.FindOne(context.TODO(), filter).Decode(&data)
	return
}

// Update updates user's mainDB data
func (mainDB MainDB) Update(userID string, update types.Main) string {
	if !isUpdateValid(update) {
		return messages.Str("invalidUserData")
	}

	msg := mainDB.uniqueFieldsExists(update)
	if msg != "" {
		return msg
	}

	mainDB.collection.UpdateOne(
		context.TODO(),
		types.Main{UserID: userID},
		map[string]types.Main{"$set": update},
	)
	return ""
}
