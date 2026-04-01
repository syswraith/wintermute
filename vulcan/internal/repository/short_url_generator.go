package repository

import (
	"math/big"
)

func ShorturlGenerator(id uint) string {
	id64 := int64(id) + 3844
	return big.NewInt(id64).Text(62)
}
