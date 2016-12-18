package main

import "fmt"

func main() {
	mcVal := MonteCarlo(500000, 1, 0.001, NaivePolicy{})
	printValueFunc(mcVal)
}

func printValueFunc(val map[Observable]float64) {
	fmt.Print("   ")
	for i := 1; i <= 10; i++ {
		if i < 10 {
			fmt.Printf("   %d  ", i)
		} else {
			fmt.Printf("  %d  ", i)
		}
	}
	fmt.Println()
	for row := 1; row <= 21; row++ {
		if row < 10 {
			fmt.Printf(" %d ", row)
		} else {
			fmt.Printf("%d ", row)
		}
		for col := 1; col <= 10; col++ {
			s := Observable{CurrentSum: row, DealerShowing: col}
			v := val[s]
			if v < 0 {
				fmt.Printf(" %.2f", v)
			} else {
				fmt.Printf("  %.2f", v)
			}
		}
		fmt.Println()
	}
}

type NaivePolicy struct{}

func (_ NaivePolicy) Action(o Observable) Action {
	return o.CurrentSum < 21
}
