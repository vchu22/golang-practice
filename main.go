package main
import ("fmt")

func main(){
	fmt.Println("Arithmetic")
	var x int = 5
	sum := x + 3 // short-hand syntax
	if x > 3 {
		fmt.Println(sum) // should print
	} else {
		fmt.Println(x) // should not print
	}


	var a [3] int // declare an array of size 3 type 'int'
	a[1] = 3      // a = [0, 3, 0]
	fmt.Println(a)

	b := [3]int{2,5,4} // short-hand syntax for array declaration
	fmt.Println(b)	   // b = [2, 5, 4]

	c := []int{2,5,4} // slice type, size will be dynamic
	c = append(c, 13)
	fmt.Println(c)	  // c = [2,5,4,13]
}
