package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

// parseTraining возвращает количество шагов, вид активности и продолжительность тренировки.
func parseTraining(data string) (int, string, time.Duration, error) {
	// Разделить строку на слайс строк.
	parts := strings.Split(data, ",")

	// Проверить, чтобы длина слайса была равна 3.
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid data format (expected 3 parts, got %d)", len(parts))
	}

	// Преобразовать первый элемент слайса (количество шагов) в тип int.
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps format: %v", err)
	}

	// Преобразовать третий элемент слайса в time.Duration.
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration format: %v", err)
	}

	// Вернуть количество шагов, вид активности, продолжительность и nil (для ошибки).
	return steps, parts[1], duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	//	Проверить, что количество steps больше 0. Если это не так, вернуть 0.
	if steps <= 0 {
		return 0
	}
	// Рассчитать дистанцию в километрах.
	return float64(steps) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
// duration time.Duration — продолжительность тренировки.
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration <= 0 {
		return 0
	}
	// Рассчитать дистанцию в километрах.
	distanceKM := distance(steps)
	// Рассчитать среднюю скорость.
	return distanceKM / (duration.Hours())
}

// TrainingInfo возвращает строку с информацией о тренировке.
// data string — данные о тренировке.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// Парсим данные о тренировке.
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		return fmt.Sprintf("Ошибка обработки данных: %v\n", err)
	}
	// Проверяем тип тренировки и возвращаем информацию о ней.
	switch activity {
	case "Бег":
		distance := distance(steps)
		speed := meanSpeed(steps, duration)
		calories := RunningSpentCalories(steps, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f \n", activity, duration.Hours(), distance, speed, calories)
	case "Ходьба":
		distance := distance(steps)
		speed := meanSpeed(steps, duration)
		calories := WalkingSpentCalories(steps, weight, height, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f \n", activity, duration.Hours(), distance, speed, calories)
	default:
		return "неизвестный тип тренировки \n"

	}
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных калорий при беге.

func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// Проверить, чтобы количество шагов steps, продолжительность duration и вес weight были больше 0.
	if steps <= 0 || duration <= 0 || weight <= 0 {
		return 0
	}
	avarageSpeed := meanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier * avarageSpeed) - runningCaloriesMeanSpeedShift) * weight

}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0
	}
	averageSpeed := meanSpeed(steps, duration)
	calories := ((walkingCaloriesWeightMultiplier * weight) +
		((averageSpeed*averageSpeed)/height)*walkingSpeedHeightMultiplier) *
		duration.Hours() * minInH
	return calories
}
