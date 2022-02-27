package util

import (
	"crypto/sha1"
	"github.com/gobuffalo/packr/v2/file/resolver/encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

func GenerateID() string {
	h := sha1.New()
	rand.Seed(time.Now().UnixNano())
	t := rand.Float64() * float64(time.Now().UnixNano())

	h.Write([]byte(strconv.Itoa(int(t))))
	return hex.EncodeToString(h.Sum(nil))[:24]
}
