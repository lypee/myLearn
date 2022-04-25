package main

type Person struct {
}

func main() {
	enter()
}

func enter() {
	testMap()
	testPerson()
	testPerson2()
	testSlice()
	testPersonSlice()
}

func testMap() map[int]int {
	res := map[int]int{}
	return res
}

func testPerson() map[int]Person{
	res := map[int]Person{}
	return res
}

func testPerson2() map[int]*Person{
	res := map[int]*Person{}
	return res
}

func testSlice()[]int{
	return []int{}
}

func testPersonSlice()[]Person{
	return []Person{}
}