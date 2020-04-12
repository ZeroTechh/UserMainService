package mainDB

import (
	"testing"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/utils"
)

func TestMainDB(t *testing.T) {
	assert := assert.New(t)
	mainDB := New()

	// Testing Create function
	mockData, _ := utils.GetMockUserData()
	mockData.UserID = mainDB.GenerateID() // Tested GenerateID
	msg := mainDB.Create(mockData)
	assert.Zero(msg)

	// Testing that Create returns invalid data message for invalid data
	assert.NotZero(mainDB.Create(types.Main{}))

	// Testing that create returns messages for already existing email and username
	mockData2, _ := utils.GetMockUserData()
	mockData2.Username = mockData.Username
	assert.NotZero(mainDB.Create(mockData2))

	mockData2, _ = utils.GetMockUserData()
	mockData2.Email = mockData.Email
	assert.NotZero(mainDB.Create(mockData2))

	// Testing Get
	returnedData := mainDB.Get(types.Main{UserID: mockData.UserID})
	assert.Equal(mockData, returnedData)

	// Testing Update
	mockData2, _ = utils.GetMockUserData()
	update := types.Main{Username: mockData2.Username}
	msg = mainDB.Update(mockData.UserID, update)
	assert.Zero(msg)

	returnedData = mainDB.Get(types.Main{UserID: mockData.UserID})
	assert.Equal(mockData2.Username, returnedData.Username)

	// Testing Update returns message for invalid update
	update = types.Main{UserID: "NN"}
	msg = mainDB.Update(mockData.UserID, update)
	assert.NotZero(msg)

	// Testing Update returns message for existing unique fields
	update = types.Main{Username: mockData2.Username}
	msg = mainDB.Update(mockData.UserID, update)
	assert.NotZero(msg)
}
