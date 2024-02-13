package bruteforce

import (
	"git.xdol.org/xdol/hashdog/internal/bruteforce/charset"
	"git.xdol.org/xdol/hashdog/internal/bruteforce/rainbow"
	"git.xdol.org/xdol/hashdog/internal/bruteforce/wordlist"
)

type Config struct {
	Length   uint8
	Threads  uint8
	Rainbow  rainbow.Config
	Wordlist wordlist.Config
	Charset  charset.Config
}
