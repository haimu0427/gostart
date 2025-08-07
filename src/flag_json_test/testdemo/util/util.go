package main

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}

	return res
}
