package utils

import (
	"blackpay_gateway/config"
	"crypto/md5"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {

	// Combine timestamp and secret key
	combinedInput := time.Now().String() + config.GetEnv("UUID_SECRET")

	// Create MD5 hash of the combined input
	md5Hash := md5.New()
	md5Hash.Write([]byte(combinedInput))
	hashBytes := md5Hash.Sum(nil)

	// Create UUID from the first 16 bytes of the hash
	uuidBytes := hashBytes[:16]

	// Modify the UUID to conform to RFC 4122 version 4
	uuidBytes[6] = (uuidBytes[6] & 0x0f) | 0x40 // version 4
	uuidBytes[8] = (uuidBytes[8] & 0x3f) | 0x80 // variant 10

	generatedUUID, _ := uuid.FromBytes(uuidBytes)
	newUUID := strings.ReplaceAll(generatedUUID.String(), "-", "")[:15]

	return newUUID
}
