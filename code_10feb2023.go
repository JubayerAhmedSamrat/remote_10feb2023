package main

import "fmt"


func add(x, y int)  (r int){

	r = x + y
	
	return 
}

func rectangle(l, b int) (area int, parameter int) {

	parameter = 2 * (l + b)
	area = l * b
	return
}


func update(a *int, t *string) {

	*a = *a + 5
	*t = *t + "Ahmed"
	return
}
//anonymous function



func main(){


	

	number := 33
	name := "jubayer "
	update(&number, &name)
	fmt.Println(number, name)

n := func(x, y int) (r int){
	r = x * y
	return

}(33, 11)

	fmt.Println(n)


	x := add(20, 12)
	a, p := rectangle(10, 10)
	fmt.Println(x, a, p)
}