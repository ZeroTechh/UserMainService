package handler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserMainService"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/ZeroTechh/UserMainService/core/mainDB"
	"github.com/ZeroTechh/UserMainService/core/types"
)

func toProto(data types.Main) *proto.Data {
	request := proto.Data{}
	copier.Copy(&request, &data)
	return &request
}

func toData(request *proto.Data) (data types.Main) {
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
	main *mainDB.MainDB
}

// init is used to initialize
func (h *Handler) init() {
	h.main = mainDB.New()
}

// Auth is used to handle Auth function
func (h Handler) Auth(ctx context.Context, request *proto.AuthRequest) (*proto.AuthResponse, error) {
	filter := types.Main{Username: request.Username, Email: request.Email}
	data, _ := h.main.Get(ctx, filter)
	valid := (data.Password == request.Password && data != types.Main{})
	// TODO Add hashing
	return &proto.AuthResponse{Valid: valid, UserID: data.UserID}, nil
}

// Add is used to handle Add function
func (h Handler) Add(ctx context.Context, request *proto.Data) (*proto.AddResponse, error) {
	data := toData(request)
	userID := h.main.ID(ctx)
	data.UserID = userID
	msg, err := h.main.Create(ctx, data)
	err = errors.Wrap(err, "Error while creating new user main data")
	return &proto.AddResponse{Message: msg, UserID: userID}, err
}

// Get is used to handle Get function
func (h Handler) Get(ctx context.Context, request *proto.GetRequest) (*proto.Data, error) {
	filter := types.Main{
		UserID:   request.UserID,
		Username: request.Username,
		Email:    request.Email,
	}
	data, err := h.main.Get(ctx, filter)
	response := toProto(data)
	return response, errors.Wrap(err, "Error while getting user main data")
}

// Update is used to handler Update function
func (h Handler) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	update := toData(request.Update)
	msg, err := h.main.Update(ctx, request.UserID, update)
	err = errors.Wrap(err, "Error while updating user main data")
	return &proto.UpdateResponse{Message: msg}, err
}

// Validate is used to handle Validate function
func (handler Handler) Validate(ctx context.Context, request *proto.Data) (*proto.ValidateResponse, error) {
	valid := mainDB.Valid(toData(request))
	return &proto.ValidateResponse{Valid: valid}, nil
}
