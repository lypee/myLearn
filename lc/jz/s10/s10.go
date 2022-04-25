package main

func fib(n int) int {
	if n < 1 {
		return 0
	}
	if n < 3 {
		return 1
	}
	arr := make([]int, n+1)
	arr[0], arr[1], arr[2] = 0, 1, 1
	for i := 3 ; i <= n ; i++{
		arr[i] = arr[i-1] + arr[i-2]
		if arr[i] >(1e9+7){
			arr[i] %= (1e9+7)
		}
	}
	return arr[n]
}
