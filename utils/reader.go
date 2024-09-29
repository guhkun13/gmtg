package utils

import (
	"os"

	"github.com/rs/zerolog/log"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) ReadFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open input file")
	}
	defer file.Close()

	return file
}
