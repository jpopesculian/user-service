package main

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port uint64
}

func NewConfig() Config {
	default_port, err := strconv.ParseUint(os.Getenv("PORT"), 0, 16)
	if err != nil {
		default_port = 5000
	}
	port := flag.Uint64("p", default_port, "Port to serve on")
	flag.Parse()

	return Config{
		*port,
	}
}
