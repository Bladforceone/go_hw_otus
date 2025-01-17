package loganalyze

import (
	"bufio"
	"fmt"
	"io"
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

func Print(w io.Writer, stats []string) error {
	if w == nil {
		return fmt.Errorf("writer is nil")
	}

	for _, stat := range stats {
		if _, err := fmt.Fprintln(w, stat); err != nil {
			return fmt.Errorf("error writing: %w", err)
		}
	}
	return nil
}
