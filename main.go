package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/libs"
	"github.com/guhkun13/gmtg/service"
	log "github.com/rs/zerolog/log"
)

func main() {
	config.InitLogger()

	regexps := config.InitRegexps()
	services := service.NewServices(regexps)

	libs.EmptyFileContent(libs.FileOutputName)

	fileInput, err := os.Open(libs.FileInputName)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open input file")
	}
	defer fileInput.Close()
	scanner := bufio.NewScanner(fileInput)

	st := time.Now()
	for scanner.Scan() {
		line := scanner.Text()
		log.Info().Str("val", line).Msg("Line")
		services.EvaluateText(line)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal().Err(err).Msg("scanner failed")
	}

	et := time.Now()
	fmt.Println("[Total execution time] : ", et.Sub(st))
}
