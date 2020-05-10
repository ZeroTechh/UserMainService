package mainDB

import (
	"reflect"

	"github.com/ZeroTechh/blaze"
	"github.com/ZeroTechh/sentinal/v2"

	"go.uber.org/zap"
)

var schemaPaths = []string{"schemas", "../core/mainDB/schemas", "core/mainDB/schemas"}

const schemaFile = "schema.yaml"

// Valid checks if data is valid
func Valid(data interface{}) bool {
	funcLog := blaze.NewFuncLog("Valid", log, zap.Any("Data", data))
	funcLog.Started()
	valid, msg, _ := sentinal.ValidateWithYAML(
		data,
		schemaFile,
		schemaPaths,
	)
	funcLog.Completed(zap.Any("Message", msg))
	return valid
}

// updateValid checks if the update is valid
func updateValid(update interface{}) bool {
	funcLog := blaze.NewFuncLog("updateValid", log, zap.Any("Update", update))
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
