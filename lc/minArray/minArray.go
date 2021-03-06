package main

func main(){

}


func minArray(numbers []int) int {
	left , right := 0 , len(numbers)  -1
	for left < right {
		mid := (left + right) / 2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] < numbers[right] {
			left = mid
		} else {
			right--
		}
	}
	return numbers[left]
}