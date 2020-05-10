package mainDB

import (
	"context"

	"github.com/ZeroTechh/VelocityCore/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZeroTechh/UserMainService/core/types"
)

// New returns a new mainDB handler struct
func New() *MainDB {
	mainDB := MainDB{}
	mainDB.init()
	return &mainDB
}

// MainDB handles user main data
type MainDB struct {
	coll *mongo.Collection
}

// init initializes client and database
func (m *MainDB) init() {
	c := utils.CreateMongoDB(dbConfig.Str("address"), log)
	db := c.Database(dbConfig.Str("db"))
	m.coll = db.Collection(dbConfig.Str("collection"))
}

// exists checks if user with certain field exists
func (m MainDB) exists(ctx context.Context, filter types.Main) bool {
	data, _ := m.Get(ctx, filter)
	return data != types.Main{}
}

// uniqueFieldsExists checks if unique fields such as username exist
func (m MainDB) uniqueFieldsExists(ctx context.Context, data types.Main) string {
	usernameExists := m.exists(ctx, types.Main{Username: data.Username})
	emailExists := m.exists(ctx, types.Main{Email: data.Email})
	if usernameExists && data.Username != "" {
		return messages.Str("usernameExists")
	} else if emailExists && data.Email != "" {
		return messages.Str("emailExists")
	}
	return ""
}

// ID generates an user id
func (m MainDB) ID(ctx context.Context) string {
	exists := true
	var userID uuid.UUID
	for exists {
		userID, _ = uuid.NewRandom()
		exists = m.exists(ctx, types.Main{UserID: userID.String()})
	}
	return userID.String()
}

// Create adds new user main data
func (m MainDB) Create(ctx context.Context, data types.Main) (string, error) {
	if !Valid(data) {
		return messages.Str("invalidUserData"), nil
	}

	msg := m.uniqueFieldsExists(ctx, data)
	if msg != "" {
		return msg, nil
	}

	_, err := m.coll.InsertOne(ctx, data)
	return "", errors.Wrap(err, "Error while inserting into db")
}

// Get returns user main data
func (m MainDB) Get(
	ctx context.Context, filter types.Main) (data types.Main, err error) {
	err = m.coll.FindOne(ctx, filter).Decode(&data)
	err = errors.Wrap(err, "Error while finding from db")
	return
}

// Update updates user's mainDB data
func (m MainDB) Update(
	ctx context.Context, userID string, update types.Main) (string, error) {
	if !updateValid(update) {
		return messages.Str("invalidUserData"), nil
	}

	msg := m.uniqueFieldsExists(ctx, update)
	if msg != "" {
		return msg, nil
	}

	_, err := m.coll.UpdateOne(
		ctx,
		types.Main{UserID: userID},
		map[string]types.Main{"$set": update},
	)
	return "", errors.Wrap(err, "Error while updating in db")
}
