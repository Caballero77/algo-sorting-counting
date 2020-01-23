package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Get("/counting", func(ctx iris.Context) {
		array := parseArray(ctx.URLParam("array"))
		sorted := sort(array)
		ctx.WriteString(arrayToString(sorted))
	})

	app.Run(iris.Addr(":80"))
}

func parseArray(array string) []int {
	numbers := strings.Split(array, ",")
	length := len(numbers)
	result := make([]int, length)
	for i := 0; i < length; i++ {
		value, _ := strconv.Atoi(numbers[i])
		result[i] = value
	}
	return result
}

func arrayToString(array []int) string {
	str := strconv.Itoa(array[0])
	length := len(array)
	for i := 1; i < length; i++ {
		str += "," + strconv.Itoa(array[i])
	}
	return str
}

func sort(array []int) []int {
	length := len(array)
	max := -2147483648
	for i := 0; i < length; i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	max = maxInt(max, length) + 1
	counts := make([]int, max)
	result := make([]int, length)

	for i := 0; i < length; i++ {
		counts[array[i]]++
	}

	sum := 0
	for i := 0; i < max; i++ {
		sum += counts[i]
		counts[i] = sum
	}

	fmt.Println(arrayToString(counts))

	for i := 0; i < length; i++ {
		result[counts[array[i]]-1] = array[i]
		counts[array[i]]--
	}

	return result
}

func maxInt(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
