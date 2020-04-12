package serviceHandler

import (
	"context"
	"testing"

	"github.com/ZeroTechh/UserService/core/types"
	proto "github.com/ZeroTechh/VelocityCore/proto/UserMainService"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/utils"
)

func TestMainDB(t *testing.T) {
	assert := assert.New(t)
	handler := New()
	ctx := context.TODO()

	// Testing Add function
	mockData, _ := utils.GetMockUserData()
	addResponse, err := handler.Add(ctx, dataToProto(mockData))
	assert.NotZero(addResponse.UserID)
	assert.Zero(addResponse.Message)
	assert.NoError(err)
	mockData.UserID = addResponse.UserID

	// Testing Auth
	authResponse, err := handler.Auth(ctx, &proto.AuthRequest{
		Username: mockData.Username,
		Password: mockData.Password,
	})
	assert.NoError(err)
	assert.True(authResponse.Valid)

	// Testing Get
	getResponse, err := handler.Get(
		ctx, &proto.GetRequest{UserID: mockData.UserID})
	assert.Equal(mockData, protoToData(getResponse))
	assert.NoError(err)

	// Testing Update
	mockData2, _ := utils.GetMockUserData()
	update := types.Main{Username: mockData2.Username}
	updateRequest := proto.UpdateRequest{
		UserID: mockData.UserID, Update: dataToProto(update),
	}
	updateResponse, err := handler.Update(ctx, &updateRequest)
	assert.NoError(err)
	assert.Zero(updateResponse.Message)
}
