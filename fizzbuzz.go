// Average of numbers between number 1-100, divided by either 3 or 5
package main

import "fmt"

func main() {
	count, total := 0, 0
	for n := 1; n <= 100; n++ {
		if n%3 == 0 || n%5 == 0 {
			count++
			total += n
		}
	}
	fmt.Println(float64(total) / float64(count))
}
