// Написать функцию находящую максимальный элемент в слайсе
// с произвольными элементами ([]interface{}) с использованием
// пользовательской функции-компаратора.
package main

import ("fmt"
"strings")

func maxElement(element ...[]string){
	for _,val := range element {
		strings.Split(val, " ")
		fmt.Println(len(val))
	}
	fmt.Println(element)
}

func main(){
	// fmt.Println("Hell")
	test := []string{"dsf dsfs dfsf sf dfs dfsd dfdsf", "dsfdsf"}

	maxElement(test)

}
