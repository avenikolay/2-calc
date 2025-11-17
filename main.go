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

func calculate(operation string, nums []int64) int64 {

	if operation == "MED" {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
	}

	result := int64(0)
	switch operation {
	case "SUM":
		for _, value := range nums {
			result += value
		}
	case "MED":
		if len(nums)%2 == 0 {
			result = (nums[len(nums)/2-1] + nums[len(nums)/2]) / 2
		} else {
			result = nums[len(nums)/2]
		}
	case "AVG":
		for _, value := range nums {
			result += value
		}
		result /= int64(len(nums))
	}
	return result
}

func main() {
	operation := getOperation()
	nums := getNums()
	result := calculate(operation, nums)
	fmt.Println(result)
}
