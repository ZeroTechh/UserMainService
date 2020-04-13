package types

// Main is used to store user main data
type Main struct {
	UserID   string `bson:"_id,omitempty,-"`
	Username string `bson:"Username,omitempty,-"`
	Email    string `bson:"Email,omitempty,-"`
	Password string `bson:"Password,omitempty,-"`
}
