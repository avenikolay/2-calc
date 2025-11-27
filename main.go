package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getOperation() string {
	var operation string
	fmt.Println("Введите операцию (AVG / SUM / MED)")

	for {
		fmt.Scan(&operation)
		if operation != "AVG" && operation != "SUM" && operation != "MED" {
			fmt.Println("Недопустимый вид операции. Допустимые значения: AVG / SUM / MED")
			continue
		}
		break
	}

	return operation
}

func getNums() []int64 {
	var userString string
	fmt.Println("Укажите список чисел, разделяя их запятой")
	fmt.Scanln(&userString)

	parts := strings.Split(userString, ",")
	numbers := make([]int64, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			fmt.Printf("Ошибка: '%s' не является числом\n", part)
			continue
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func getMedian(nums []int64) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if len(nums)%2 == 0 {
		return (nums[len(nums)/2-1] + nums[len(nums)/2]) / 2
	}
	return nums[len(nums)/2]
}

func getSum(nums []int64) int64 {
	result := int64(0)
	for _, value := range nums {
		result += value
	}
	return result
}
func getAvg(nums []int64) int64 {
	result := int64(0)
	for _, value := range nums {
		result += value
	}
	result /= int64(len(nums))
	return result
}

var actions = map[string]func([]int64) int64{
	"MED": getMedian,
	"SUM": getSum,
	"AVG": getAvg,
}

func main() {
	operation := getOperation()
	nums := getNums()
	calculateFunc := actions[operation]
	result := calculateFunc(nums)
	fmt.Println(result)
}
