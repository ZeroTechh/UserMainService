package handler

import (
	"context"
	"testing"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserMainService"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserMainService/core/types"
	"github.com/ZeroTechh/UserMainService/core/utils"
)

func TestServiceHandler(t *testing.T) {
	assert := assert.New(t)
	h := New()
	ctx := context.TODO()

	// Testing Add function
	data := utils.Mock()
	addResponse, err := h.Add(ctx, toProto(data))
	assert.NotZero(addResponse.UserID)
	assert.Zero(addResponse.Message)
	assert.NoError(err)
	data.UserID = addResponse.UserID

	// Testing Auth
	authResponse, err := h.Auth(ctx, &proto.AuthRequest{
		Username: data.Username,
		Password: data.Password,
	})
	assert.NoError(err)
	assert.True(authResponse.Valid)

	// Testing Get
	getResponse, err := h.Get(
		ctx, &proto.GetRequest{UserID: data.UserID})
	assert.Equal(data, toData(getResponse))
	assert.NoError(err)

	// Testing Update
	data2 := utils.Mock()
	update := types.Main{Username: data2.Username}
	updateRequest := proto.UpdateRequest{
		UserID: data.UserID, Update: toProto(update),
	}
	updateResponse, err := h.Update(ctx, &updateRequest)
	assert.NoError(err)
	assert.Zero(updateResponse.Message)

	// Testing Validate
	data = utils.Mock()
	valid, err := h.Validate(ctx, toProto(data))
	assert.True(valid.Valid)
	assert.NoError(err)
}
