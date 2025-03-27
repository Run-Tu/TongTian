package main

import "fmt"

func g(arg any) {
	a, _ := arg.(string)
	fmt.Println(a)
}

func main() {
	a := 3.14
	g(a)

	m := map[int]bool{}
	value := m[1]
	fmt.Printf("%t\n", value)

	ch := make(chan int, 1)
	close(ch)
	v := <-ch
	fmt.Printf("%d\n", v)
}
