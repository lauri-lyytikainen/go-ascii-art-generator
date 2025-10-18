package parameters

import (
	"flag"
)

type Config struct {
	InputFile  string
	OutputFile string
}

func ParseParameters() (*Config, error) {

	cfg := &Config{}

	flag.StringVar(&cfg.InputFile, "i", "", "input file")
	flag.StringVar(&cfg.OutputFile, "o", "", "output file")
	flag.Parse()

	// TODO: Decide if error is needed
	return cfg, nil
}
