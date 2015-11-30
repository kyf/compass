package data

import (
	"log"
)

var (
	logger *log.Logger
)

func initLogger(_logger *log.Logger) {
	logger = _logger
}
