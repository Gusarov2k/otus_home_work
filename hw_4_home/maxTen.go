// Написать функцию, которая получает на вход текст и возвращает 10 самых часто встречающихся слов без учета словоформ
package main

import (
	"fmt"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func unique(element string) []kv {
	stringArr := strings.Split(element, " ")
	keys := make(map[string]int)

	for _, entry := range stringArr {
		if _, value := keys[entry]; !value {
			keys[entry] = 0
		}
	}

	for key := range keys {
		for _, v := range stringArr {
			if key == v {
				keys[key] += 1
			}
		}
	}

	var ss []kv
	for k, v := range keys {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	return ss
}

func main() {
	test := "man man dsfs dfsf sf dfs dfsd man dfdsf d d d d"

	fmt.Println(unique(test))
}
