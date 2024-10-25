package tlog

import (
	"io"
	"os"
)

type Handler struct {
	level Level
	w     []io.Writer
}

func NewHandler(level Level, filePath string) *Handler {

	writers := []io.Writer{
		os.Stdout,
	}

	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic("cannot open log file: " + err.Error())
		}

		writers = append(writers, file)
	}

	return &Handler{
		level: level,
		w:     writers,
	}
}
