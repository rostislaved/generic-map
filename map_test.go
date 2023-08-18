package gmap

import (
	"fmt"
	"strconv"
	"testing"
)

var n = 100000

var kLookup = n - n/2

func BenchmarkAccess1mapG(b *testing.B) {
	numbetOfElementsList := []int{10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

	for _, numberOfElements := range numbetOfElementsList {
		m := New[int, int](numberOfElements)

		for i := 0; i < n; i++ {
			m.Assign(i, i+1)
		}

		b.Run(strconv.Itoa(numberOfElements), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = m.Access1(kLookup)
			}
		})

		fmt.Printf("LoadFactor: %1.3f\n", m.LoadFactor())
	}
}

func BenchmarkAccess1map(b *testing.B) {
	m := make(map[int]int, n)

	for i := 0; i < n; i++ {
		m[i] = i + 1
	}

	for i := 0; i < b.N; i++ {
		_ = m[kLookup]
	}
}

func ExampleHmap_Access1() {
	m := New[int, int](5)

	m.Assign(1, 11)
	m.Assign(2, 22)

	fmt.Println(m.Access1(1))
	fmt.Println(m.Access1(3))
}
