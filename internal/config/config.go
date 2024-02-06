package config

import (
	"git.xdol.org/xdol/hashdog/internal/bruteforce"
	"git.xdol.org/xdol/hashdog/internal/log"
)

type Config struct {
	bruteforce bruteforce.Config
	length     uint8
	threads    uint8
	log        log.Config
}
