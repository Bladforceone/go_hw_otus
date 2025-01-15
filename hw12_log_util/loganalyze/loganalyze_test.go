package loganalyze

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnalyze(t *testing.T) {
	content := `2023-10-01 12:00:00 [INFO] User logged in
2023-10-01 12:01:00 [ERROR] Failed to connect
2023-10-01 12:02:00 [INFO] User logged out
2023-10-01 12:03:00 [WARNING] Disk space low`

	tmpFile, err := os.CreateTemp("", "test_log_*.log")
	require.NoError(t, err, "Ошибка при создании временного файла")
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	require.NoError(t, err, "Ошибка при записи во временный файл")
	tmpFile.Close()

	tests := []struct {
		name     string
		level    string
		expected []string
	}{
		{
			name:  "Анализ логов уровня INFO",
			level: "INFO",
			expected: []string{
				"2023-10-01 12:00:00 [info] user logged in",
				"2023-10-01 12:02:00 [info] user logged out",
			},
		},
		{
			name:  "Анализ логов уровня ERROR",
			level: "ERROR",
			expected: []string{
				"2023-10-01 12:01:00 [error] failed to connect",
			},
		},
		{
			name:     "Анализ логов с несуществующим уровнем",
			level:    "DEBUG",
			expected: []string(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := Analyze(tmpFile.Name(), tt.level)
			require.NoError(t, err, "Ошибка при анализе лог-файла")
			assert.Equal(t, tt.expected, stats, "Некорректный результат анализа логов")
		})
	}
}

func TestPrint(t *testing.T) {
	tests := []struct {
		name     string
		stats    []string
		expected string
	}{
		{
			name: "Запись логов в файл",
			stats: []string{
				"2023-10-01 12:00:00 [INFO] User logged in",
				"2023-10-01 12:02:00 [INFO] User logged out",
			},
			expected: `2023-10-01 12:00:00 [INFO] User logged in
2023-10-01 12:02:00 [INFO] User logged out
`,
		},
		{
			name:     "Запись пустого списка логов",
			stats:    []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем временный файл для записи
			tmpFile, err := os.CreateTemp("", "test_output_*.log")
			require.NoError(t, err, "Ошибка при создании временного файла")
			defer os.Remove(tmpFile.Name())

			// Записываем данные в файл
			err = Print(tt.stats, tmpFile.Name())
			require.NoError(t, err, "Ошибка при записи статистики в файл")

			// Читаем файл и проверяем его содержимое
			fileContent, err := os.ReadFile(tmpFile.Name())
			require.NoError(t, err, "Ошибка при чтении файла")

			assert.Equal(t, tt.expected, string(fileContent), "Некорректное содержимое файла")
		})
	}
}
