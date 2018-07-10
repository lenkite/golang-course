package main

import "fmt"

func sumProdDiff(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func sumProdDiffNamed(a, b int) (sum, prod, diff int) { // named results are auto-documenting
	//named results are zero valued before assignment
	fmt.Printf("(inside fn) sum before assignment = %d\n", sum)
	sum = a + b
	prod = a * b
	diff = a - b
	return //bare return used when results are named
}

func main() {
	s, p, d := sumProdDiff(1, 2)
	fmt.Println("a>", s, p, d)

	s, p, _ = sumProdDiffNamed(3, 4)
	fmt.Println("b>", s, p, d)
}
