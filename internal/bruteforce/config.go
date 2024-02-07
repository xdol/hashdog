package bruteforce

import (
	"git.xdol.org/xdol/hashdog/internal/bruteforce/charset"
	"git.xdol.org/xdol/hashdog/internal/bruteforce/rainbow"
)

type Config struct {
	Length   uint8
	Threads  uint8
	Rainbow  rainbow.Config
	Wordlist rainbow.Config
	Charset  charset.Config
}
