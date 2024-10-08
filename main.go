package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/guhkun13/gmtg/config"
	"github.com/guhkun13/gmtg/service"
	"github.com/guhkun13/gmtg/utils"
	log "github.com/rs/zerolog/log"
)

// set isDebug to false by default so that DEBUG message not showed

const FileInputName string = "input.txt"
const FileOutputName string = "output.txt"

func main() {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
		panic(err)
	}
	config.InitLogger(env)

	regexps := config.InitRegexps()
	services := service.NewServices(regexps)

	utils.EmptyFileContent(FileOutputName)

	fileInput, err := os.Open(FileInputName)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open input file")
	}
	defer fileInput.Close()
	scanner := bufio.NewScanner(fileInput)

	st := time.Now()
	for scanner.Scan() {
		text := scanner.Text()
		answer := services.EvaluateText(text)

		log.Info().
			Str("0-question", text).
			Str("1-answer", answer).
			Msg("result")

		if answer != "" {
			utils.WriteToFileOutput(FileOutputName, answer)
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal().Err(err).Msg("scanner failed")
	}

	et := time.Now()
	fmt.Println("[Total execution time] : ", et.Sub(st))
}
