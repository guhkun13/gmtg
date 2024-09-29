package utils

import (
	"os"

	"github.com/guhkun13/gmtg/libs"
	"github.com/rs/zerolog/log"
)

func WriteToOutput(content string) {
	// log.Debug().Str("content", content).Msg("WriteToOutput")

	// open output file
	fileOutput, err := os.OpenFile(libs.FileOutputName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open file")
	}
	defer fileOutput.Close()

	// Write the content to the file
	if _, err := fileOutput.WriteString(content + " \n"); err != nil {
		log.Fatal().Err(err).Msg("Failed to write to file")
	}
}
