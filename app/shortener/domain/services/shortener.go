package services

import (
	"fmt"

	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
	"math/big"
)

func ShortInput(input string, length int) string {
	sha := utils.CreateSha256(input)
	n := new(big.Int).SetBytes(sha).Uint64()
	encode := utils.CreateBase58([]byte(fmt.Sprintf("%d", n)))
	return encode[:length]
}
