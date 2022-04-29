package randomizer

import (
    "crypto/rand"
)

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func Randomize(length int) string{
	ll := len(chars)
    b := make([]byte, length)
    rand.Read(b) // generates len(b) random bytes
    for i := 0; i < 20; i++ {
        b[i] = chars[int(b[i])%ll]
    }
    return string(b)
}