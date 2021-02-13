package main

import "fmt"

func main() {

	// sales := [][]int{
	// 	{2000, 4000, 6000, 700},
	// 	{4500, 560},
	// 	{340, 900, 800},
	// }
	var sales = make([][]int, 0)
	sales = append(sales, []int{2000, 4000, 6000, 700})
	sales = append(sales, []int{4500, 560})
	sales = append(sales, []int{340, 900, 800})
	fmt.Println(sales)

	for i, v := range sales {
		sum := 0
		for _, p := range v {
			sum += p
		}
		fmt.Printf("Day %d Sales: %d\n", i+1, sum)
	}
}

