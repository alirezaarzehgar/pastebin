package factory

import (
	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/config/dotenv"
)

type ConfigLoader uint

const (
	ConfigLoaderDotEnv ConfigLoader = iota
)

func New(cl ConfigLoader) config.ConfigLoader {
	switch cl {
	case ConfigLoaderDotEnv:
		return dotenv.New()
	default:
		panic("invalid ConfigLoader inserted")
	}
}
