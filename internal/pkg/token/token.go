package token

import (
	"github.com/simonnik/GB_Backend1_CW_GO/internal/pkg/random"
	"github.com/speps/go-hashids/v2"
)

// GenerateToken generated random hash
// Usage:
// arr := token.GenerateToken()
//
// Output: agQIBt7HbD
func GenerateToken(l int, salt string) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = l
	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64(random.RangeInt(0, 100, 4))

	return e
}
