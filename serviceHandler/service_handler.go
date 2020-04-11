package serviceHandler

import (
	"context"

	"github.com/ZeroTechh/UserService/core/types"
	proto "github.com/ZeroTechh/VelocityCore/proto/UserMainService"
	"github.com/jinzhu/copier"

	"github.com/ZeroTechh/UserMainService/core/mainDB"
)

// New returns a new service handler
func New() *Handler {
	handler := Handler{}
	handler.init()
	return &handler
}

// Handler is used to handle all user main service functions
type Handler struct {
	mainDB *mainDB.MainDB
}

// Init is used to initialize
func (handler *Handler) init() {
	handler.mainDB = mainDB.New()
}

// Auth is used to handle Auth function
func (handler Handler) Auth(ctx context.Context, request *proto.AuthRequest) (*proto.AuthResponse, error) {
	filter := types.Main{Username: request.Username, Email: request.Email}
	data := handler.mainDB.Get(filter)
	valid := data.Password == request.Password && data != types.Main{}
	// TODO Add hashing
	return &proto.AuthResponse{Valid: valid}, nil
}

// Add is used to handle Add function
func (handler Handler) Add(ctx context.Context, request *proto.Data) (*proto.AddResponse, error) {
	var data types.Main
	copier.Copy(&data, &request)

	userID := handler.mainDB.GenerateID()
	msg := handler.mainDB.Create(data)

	return &proto.AddResponse{Message: msg, UserID: userID}, nil
}

// Get is used to handle Get function
func (handler Handler) Get(ctx context.Context, request *proto.GetRequest) (response *proto.Data, err error) {
	filter := types.Main{
		UserID:   request.UserID,
		Username: request.Username,
		Email:    request.Email,
	}

	data := handler.mainDB.Get(filter)
	copier.Copy(&response, &data)

	return
}

// Update is used to handler Update function
func (handler Handler) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	var update types.Main
	copier.Copy(&update, &request.Update)

	msg := handler.mainDB.Update(request.UserID, update)

	return &proto.UpdateResponse{Message: msg}, nil
}
