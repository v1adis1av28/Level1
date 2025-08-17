package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить эти значения в группы с шагом 10 градусов.
// Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.
// Пояснение: диапазон -20 включает значения от -20 до -29.9, диапазон 10 – от 10 до 19.9, и т.д. Порядок в подмножествах не важен.

func group(temps *[]float64) map[int][]float64 {
	mp := make(map[int][]float64)

	for _, val := range *temps {
		bucket := int(val/10) * 10
		mp[bucket] = append(mp[bucket], val)
	}

	return mp
}

func main() {

	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 100.4, -239.4, 110.0, 112.0, 0.9, 3.4}
	mps := group(&temps)

	for key, values := range mps {
		fmt.Printf("Bucket %d: %v\n", key, values)
	}
}
