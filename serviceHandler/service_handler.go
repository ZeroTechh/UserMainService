package serviceHandler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserMainService"
	"github.com/jinzhu/copier"

	"github.com/ZeroTechh/UserMainService/core/mainDB"
	"github.com/ZeroTechh/UserMainService/core/types"
)

func dataToProto(data types.Main) *proto.Data {
	request := proto.Data{}
	copier.Copy(&request, &data)
	return &request
}

func protoToData(request *proto.Data) (data types.Main) {
	copier.Copy(&data, &request)
	return
}

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
	return &proto.AuthResponse{Valid: valid, UserID: data.UserID}, nil
}

// Add is used to handle Add function
func (handler Handler) Add(ctx context.Context, request *proto.Data) (*proto.AddResponse, error) {
	data := protoToData(request)

	userID := handler.mainDB.GenerateID()
	data.UserID = userID
	msg := handler.mainDB.Create(data)

	return &proto.AddResponse{Message: msg, UserID: userID}, nil
}

// Get is used to handle Get function
func (handler Handler) Get(ctx context.Context, request *proto.GetRequest) (*proto.Data, error) {
	filter := types.Main{
		UserID:   request.UserID,
		Username: request.Username,
		Email:    request.Email,
	}

	data := handler.mainDB.Get(filter)
	response := dataToProto(data)

	return response, nil
}

// Update is used to handler Update function
func (handler Handler) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	update := protoToData(request.Update)
	msg := handler.mainDB.Update(request.UserID, update)
	return &proto.UpdateResponse{Message: msg}, nil
}
