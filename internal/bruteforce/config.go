package bruteforce

import (
	"git.xdol.org/xdol/hashdog/internal/bruteforce/charset"
	"git.xdol.org/xdol/hashdog/internal/bruteforce/rainbow"
)

type Config struct {
	rainbow  rainbow.Config
	wordlist rainbow.Config
	charset  charset.Config
}
