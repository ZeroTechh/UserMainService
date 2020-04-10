package mainDB

import (
	"reflect"

	"github.com/ZeroTechh/blaze"
	"github.com/ZeroTechh/sentinal/v2"

	"go.uber.org/zap"
)

var schemaPaths = []string{"schemas"}

const schemaFile = "schema.yaml"

// checks if data is valid
func isDataValid(data interface{}) bool {
	funcLog := blaze.NewFuncLog(
		"isDataValid",
		log,
		zap.Any("Data", data),
	)
	funcLog.Started()

	valid, msg, _ := sentinal.ValidateWithYAML(
		data,
		schemaFile,
		schemaPaths,
	)

	funcLog.Completed(zap.Any("Message", msg))
	return valid
}

// checks if the update is valid
func isUpdateValid(update interface{}) bool {
	funcLog := blaze.NewFuncLog(
		"isUpdateValid",
		log,
		zap.Any("Update", update),
	)
	funcLog.Started()

	valid, msg, _ := sentinal.ValidateFieldsWithYAML(
		update,
		schemaFile,
		schemaPaths,
	)

	// checking if the update data is not trying to update UserID
	if reflect.ValueOf(update).FieldByName("UserID").String() != "" {
		valid = false
		msg = map[string][]string{
			"UserID": []string{"Tring to update UserID"},
		}
	}

	funcLog.Completed(zap.Any("Message", msg))
	return valid
}
