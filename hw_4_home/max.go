// Написать функцию, которая получает на вход текст и возвращает 10 самых часто встречающихся слов без учета словоформ
package main

import (
	"fmt"
	// "sort"
	"strings"
)

func unique(element string) []string {
	intSlice := strings.Split(element, " ")
	keys := make(map[string]int)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = 0
			list = append(list, entry)
		}
	}

	for key := range keys {
		for _, v := range intSlice {
			if key == v {
				keys[key] += 1
			}
		}
	}

	fmt.Println(keys)
	return list
}

func main() {
	test := "man man dsfs dfsf sf dfs dfsd man dfdsf"

	fmt.Println(unique(test))
}
