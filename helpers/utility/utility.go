package utility

import "crypto/rand"

var groupCodeChars = []byte("1234567890ABCDEFGHJKLMNPQRSTUVWXYZ")
var groupCodeLength = 6

func GenerateGroupCode() string {
	return generateRandomString(groupCodeLength, groupCodeChars)
}

func generateRandomString(length int, chars []byte) string {
	l := length
	clen := len(chars)
	maxrb := 255 - (256 % clen)
	pass := make([]byte, l)
	randomBytes := make([]byte, l+(l/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(randomBytes); err != nil {
			panic(err)
		}
		for _, a := range randomBytes {
			c := int(a)
			if c >= maxrb {
				continue
			}
			pass[i] = chars[c%clen]
			i++
			if i == l {
				return string(pass)
			}
		}
	}
}
