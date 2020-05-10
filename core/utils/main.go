package utils

import (
	"math/rand"
	"time"

	"github.com/ZeroTechh/UserMainService/core/types"
)

func randStr(length int) string {
	charset := "1234567890abcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Mock returns mock user main data for testing
func Mock() types.Main {
	randomStr := randStr(10)
	mockUserData := types.Main{
		Username: randomStr,
		Password: randomStr,
		Email:    randomStr + "@gmail.com",
	}
	return mockUserData
}
