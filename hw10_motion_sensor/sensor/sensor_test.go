package sensor_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Bladforceone/go_hw_otus/hw10_motion_sensor/sensor"
	"github.com/stretchr/testify/assert"
)

func TestSensData(t *testing.T) {
	// Создаем канал для передачи данных
	dataChan := make(chan int, 600)

	// Создаем слайс с заранее выделенной памятью
	data := make([]int, 0, 600)

	// Запускаем SensData в отдельной горутине
	go sensor.SensData(dataChan)

	// Собираем данные из канала
	for d := range dataChan {
		data = append(data, d)
	}

	// Проверяем, что количество данных соответствует ожидаемому
	assert.Len(t, data, 600, "Данные должны содержать 600 элементов")

	// Проверяем, что данные находятся в пределах 0-999 (рандом от 0 до 999)
	for _, d := range data {
		assert.True(t, d >= 0 && d < 1000, "Значение должно быть в диапазоне [0, 1000)")
	}
}

func TestProcessData(t *testing.T) {
	// Создаем каналы
	dataChan := make(chan int, 600)
	processChan := make(chan int, 60)

	// Создаем слайс с заранее выделенной памятью
	processedData := make([]int, 0, 60)

	// Запускаем SensData в отдельной горутине
	go sensor.SensData(dataChan)

	// Запускаем ProcessData в отдельной горутине
	go sensor.ProcessData(dataChan, processChan)

	// Собираем данные из processChan
	for pd := range processChan {
		processedData = append(processedData, pd)
	}

	// Проверяем, что данных в канале processChan ровно 60 (по 10 элементов на 1 сумму)
	assert.Len(t, processedData, 60, "Должно быть 60 сумм в канале для 600 элементов")

	// Проверяем, что каждая сумма данных является допустимой
	for _, sum := range processedData {
		assert.True(t, sum >= 0, "Сумма должна быть неотрицательной")
	}
}

func TestProcessDataWithNoData(t *testing.T) {
	// Создаем каналы
	dataChan := make(chan int)
	processChan := make(chan int)

	// Запускаем ProcessData без данных
	go sensor.ProcessData(dataChan, processChan)

	// Пытаемся получить данные из processChan, но канал должен быть закрыт сразу
	select {
	case <-processChan:
		t.Fatal("Канал processChan не должен содержать данных")
	default:
		// Нормально, если ничего не происходит
	}
}

func TestSensDataAndProcessDataInteraction(t *testing.T) {
	// Создаем каналы
	dataChan := make(chan int, 600)
	processChan := make(chan int, 60)

	// Создаем слайс с заранее выделенной памятью
	processedData := make([]int, 0, 60)

	// Запускаем SensData и ProcessData в горутинах
	go sensor.SensData(dataChan)
	go sensor.ProcessData(dataChan, processChan)

	// Собираем данные из processChan
	for pd := range processChan {
		processedData = append(processedData, pd)
	}

	// Проверяем, что данные были корректно обработаны
	assert.Len(t, processedData, 60, "Должно быть 60 сумм в канале для 600 элементов")

	// Проверяем, что суммы данных соответствуют ожидаемым диапазонам
	for _, sum := range processedData {
		assert.True(t, sum >= 0, "Сумма должна быть неотрицательной")
	}
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
