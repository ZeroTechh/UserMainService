package mainDB

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/ZeroTechh/UserMainService/core/types"
	"github.com/ZeroTechh/UserMainService/core/utils"
)

func TestMainDB(t *testing.T) {
	assert := assert.New(t)
	m := New()
	ctx := context.TODO()

	// Testing Create function
	data := utils.Mock()
	data.UserID = m.ID(ctx) // Tested ID
	msg, err := m.Create(ctx, data)
	assert.Zero(msg)
	assert.NoError(err)

	// Testing that Create returns invalid data message for invalid data
	msg, err = m.Create(ctx, types.Main{})
	assert.NotZero(msg)
	assert.NoError(err)

	// Testing that create returns messages for already existing email and username
	data2 := utils.Mock()
	data2.Username = data.Username
	msg, err = m.Create(ctx, data2)
	assert.NotZero(msg)
	assert.NoError(err)

	data2 = utils.Mock()
	data2.Email = data.Email
	msg, err = m.Create(ctx, data2)
	assert.NotZero(msg)
	assert.NoError(err)

	// Testing Get
	returnedData, err := m.Get(ctx, types.Main{UserID: data.UserID})
	assert.Equal(data, returnedData)
	assert.NoError(err)

	// Testing Update
	data2 = utils.Mock()
	update := types.Main{Username: data2.Username}
	msg, err = m.Update(ctx, data.UserID, update)
	assert.Zero(msg)
	assert.NoError(err)

	returnedData, err = m.Get(ctx, types.Main{UserID: data.UserID})
	assert.Equal(data2.Username, returnedData.Username)
	assert.NoError(err)

	// Testing Update returns message for invalid update
	update = types.Main{UserID: "NN"}
	msg, err = m.Update(ctx, data.UserID, update)
	assert.NotZero(msg)
	assert.NoError(err)

	// Testing Update returns message for existing unique fields
	update = types.Main{Username: data2.Username}
	msg, err = m.Update(ctx, data.UserID, update)
	assert.NotZero(msg)
	assert.NoError(err)
}
