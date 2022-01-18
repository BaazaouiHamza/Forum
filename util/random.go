package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "adcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomString generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[(rand.Intn(k))]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomOwner generate a random post title
func RandomPostTitle() string {
	return RandomString(6)
}

//RandomOwner generate a random post description/text
func RandomPostDescriptionOrText() string {
	return RandomString(20)
}

//RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

//RandomEmail generates a random image
func RandomImage() string {
	return fmt.Sprintf("%s.png", RandomString(6))
}

//RandomOwner generate a random owner name
func RandomUser() string {
	return RandomString(6)
}
