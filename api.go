package main

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Get("/counting", func(ctx iris.Context) {
		ctx.Write(parseAndSort([]byte("[" + ctx.URLParam("array") + "]")))
	})
	app.Post("/counting", func(ctx iris.Context) {
		body, _ := ctx.GetBody()
		ctx.Write(parseAndSort(body))
	})
	app.Run(iris.Addr(":80"))
}

func parseAndSort(bytes []byte) []byte {
	var array []int
	json.Unmarshal(bytes, &array)

	b, _ := json.Marshal(map[string][]int{"result": sort(array)})

	return b
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
