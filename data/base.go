package data

import (
	"log"
)

var (
	logger *log.Logger
)

func InitLogger(_logger *log.Logger) {
	logger = _logger
}
