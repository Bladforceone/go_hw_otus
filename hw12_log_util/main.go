package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Bladforceone/go_hw_otus/hw12_log_util/loganalyze"
)

func main() {
	fileFl := flag.String("file", "",
		"указывает путь к анализируемому лог-файлу (обязательный флаг)")
	levelFl := flag.String("level", "info",
		"указывает уровень логов для анализа (необязательный флаг)")
	outputFl := flag.String("output", "",
		"указывает путь к файлу, в который будет записана статистика (необязательный флаг)")
	flag.Parse()

	filepath := *fileFl
	if filepath == "" {
		filepath = os.Getenv("LOG_ANALYZER_FILE")
	}

	level := *levelFl
	if level == "" {
		level = os.Getenv("LOG_ANALYZER_LEVEL")
	}

	stats, err := loganalyze.Analyze(filepath, level)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := *outputFl
	if output == "" && flag.Lookup("output") != nil {
		output = os.Getenv("LOG_ANALYZER_OUTPUT")
	}

	var writer io.Writer

	if output == "" {
		writer = os.Stdout
	} else {
		file, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("error opening file:", err)
			return
		}
		defer file.Close()
		writer = file
	}

	err = loganalyze.Print(writer, stats)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
