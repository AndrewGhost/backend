package recursive

import "fmt"

func move(diskNum int, from string, to string) {
	fmt.Printf("disk %d From %s TO %s\n", diskNum, from, to)
}

func Hanoi(n int, A string, B string, C string) {
	if 1 == n {
		move(1, A, C)
	} else {
		Hanoi(n-1, A, C, B)
		move(n, A, C)
		Hanoi(n-1, B, A, C)
	}
}
