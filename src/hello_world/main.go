package main
import ("fmt")

func main(){
	var x int = 5
	sum := x + 3 // short-hand syntax
	fmt.Println("Hello World")
	if x > 3 {
		fmt.Println(sum) // should print
	} else {
		fmt.Println(x) // should not print
	}

	var a [3] int
	a[1] = 3
	fmt.Println(a)

	b := [3]int{2,5,4}
	fmt.Println(b)

	c := []int{2,5,4} // slice type, size will be dynamic
	c = append(c, 13)
	fmt.Println(c)
}