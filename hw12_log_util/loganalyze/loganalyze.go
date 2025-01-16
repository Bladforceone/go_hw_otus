package loganalyze

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Analyze(filepath, level string) ([]string, error) {
	var stats []string

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	level = strings.ToLower(level)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, level) {
			stats = append(stats, line)
		}
	}

	if errScan := scanner.Err(); errScan != nil {
		return nil, fmt.Errorf("error reading file: %w", errScan)
	}

	return stats, err
}

func Print(stats []string, output string) error {
	if output == "" {
		for _, stat := range stats {
			fmt.Println(stat)
		}
		return nil
	}

	file, errOpen := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if errOpen != nil {
		return fmt.Errorf("error opening file: %w", errOpen)
	}
	defer file.Close()

	for _, stat := range stats {
		_, err := file.WriteString(stat + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	return nil
}
