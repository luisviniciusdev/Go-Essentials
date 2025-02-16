package main

func AddGenerics[T int | float64](a, b T) T {
	return a + b
}
