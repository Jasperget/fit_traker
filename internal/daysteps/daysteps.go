package daysteps

import (
	"fit_traker/internal/spentcalories"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

// parsePackage парсит строку с данными о количестве шагов и времени.
// Возвращает количество шагов, продолжительность и ошибку.
func parsePackage(data string) (int, time.Duration, error) {
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, fmt.Errorf("invalid data the number of values ​​is not equal to 2")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid data, %s not int", slice[0])
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("invalid data 3")
	}
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, duration, nil

}

// DayActionInfo возвращает информацию о дневной активности.
// Принимает строку с данными о количестве шагов и времени, вес и рост пользователя.
// Возвращает строку с информацией о дневной активности.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка парсинга данных:", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}

	// Вычисление дистанции
	distance := float64(steps) * StepLength / 1000

	// Вычисление калорий
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	// Формирование строки результата
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
