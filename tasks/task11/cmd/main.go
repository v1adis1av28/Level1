package main

import (
	"fmt"
	"sort"
)

// Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.
// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

func main() {
	setA := []int{1, 2, 2, 3}
	setB := []int{2, 3, 4, 2}

	crossAlgo := FindIntersectionWithAlgorithm(setA, setB)
	crossMap := FindIntersectionWithMap(setA, setB)

	fmt.Println("result using algorithm:")
	fmt.Println(crossAlgo)

	fmt.Println("result using map:")
	fmt.Println(crossMap)
}

// Проходим по обоим множества и добавляем в мапу значение ключа - это значение из множества, а значение - это количество вхождений в множествах, если их больше 1 будем записывать
// в результирующее множество(затратное по памяти решение)
func FindIntersectionWithMap(setA []int, setB []int) []int {
	set := make(map[int]bool)
	resSet := make(map[int]bool)
	res := []int{}

	for _, val := range setB {
		set[val] = true
	}

	for _, val := range setA {
		if set[val] && !resSet[val] {
			res = append(res, val)
			resSet[val] = true
		}
	}

	return res
}

// Другое решение это использование двух указателей которые будут, провреять равенство или неравенство эл-ов в массиве
// Но так как это неоптимальное решение из-за неупорядоченности множеств по заданию.
// Сперва отсортируем массивы и реализуем алгоритм с использованием двух указателей
// За счет времени на сортировку не выигрываем, но по памяти лучше чем с мапой
func FindIntersectionWithAlgorithm(setA []int, setB []int) []int {
	indA := 0
	indB := 0
	res := []int{}
	sort.Ints(setA)
	sort.Ints(setB)
	for indA < len(setA) && indB < len(setB) {
		if setA[indA] == setB[indB] {
			if len(res) == 0 || res[len(res)-1] != setA[indA] {
				res = append(res, setA[indA])
			}
			indA++
			indB++
		} else if setA[indA] < setB[indB] {
			indA++
		} else {
			indB++
		}
	}
	return res
}
