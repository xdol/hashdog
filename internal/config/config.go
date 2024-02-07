package config

import (
	"git.xdol.org/xdol/hashdog/internal/bruteforce"
	"git.xdol.org/xdol/hashdog/internal/log"
)

type Config struct {
	Bruteforce bruteforce.Config
	Log        log.Config
}
