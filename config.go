package main

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port           uint64
	PrivateKeyPath string
}

var config Config

func InitConfig() {
	defaultPort, err := strconv.ParseUint(os.Getenv("PORT"), 0, 16)
	if err != nil {
		defaultPort = 5000
	}
	port := flag.Uint64("p", defaultPort, "Port to serve on")

	defaultPrivateKeyPath := os.Getenv("PRIVATE_KEY_PATH")
	if len(defaultPrivateKeyPath) == 0 {
		defaultPrivateKeyPath = "id_rsa"
	}
	privateKeyPath := flag.String("sk", defaultPrivateKeyPath, "Path to RSA Private Key")

	flag.Parse()

	config = Config{
		*port,
		*privateKeyPath,
	}
}
