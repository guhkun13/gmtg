package libs

import (
	"log"
	"os"
)

func EmptyFileContent(filename string) {
	if err := os.Truncate(filename, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
}
