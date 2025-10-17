package parameters

import (
	"flag"
)

type Config struct {
	InputFile string
}

func ParseParameters() (*Config, error) {

	cfg := &Config{}

	flag.StringVar(&cfg.InputFile, "i", "", "input file")
	flag.Parse()

	return cfg, nil
}
