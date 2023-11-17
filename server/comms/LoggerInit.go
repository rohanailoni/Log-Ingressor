package comms

import (
	"log"
)

func LoggerInit() *log.Logger {
	return log.New(
		// Log output, os.Stdout for standard output
		// You can replace it with a file or any other io.Writer
		log.Writer(),
		"INFO: ",
		// Log flags, you can adjust these as needed
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}
