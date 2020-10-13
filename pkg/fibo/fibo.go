package fibo

func Fibo(n int) int {
	if n < 3 {
		return 1
	}
	f := [2]int{1, 1}
	for i := 2; i <= n; i++ {
		f[i%2] = f[0] + f[1]
	}
	return f[n%2]
}
