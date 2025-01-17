package main

import (
	"flag"
	"fmt"
	"log"
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

	if output == "" {
		if err = loganalyze.Print(os.Stdout, stats); err != nil {
			log.Fatalf("error printing statistics: %v", err)
		}
	} else {
		file, errFile := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if errFile != nil {
			fmt.Println("error opening file:", errFile)
			os.Exit(1)
		}
		defer file.Close()

		if err = loganalyze.Print(file, stats); err != nil {
			log.Fatalf("error printing statistics: %v", err)
		}
	}
}
