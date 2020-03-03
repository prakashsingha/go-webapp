package helper

import (
	"crypto/sha512"
	"encoding/base64"
)

func Hasher(args ...string) string {
	hasher := sha512.New()

	for _, val := range args {
		hasher.Write([]byte(val))
	}
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	// fmt.Println(pwd)
	return pwd
}
