package main

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "p1",
		Age:  1,
	}

	p2 := Person{
		Name: "p2",
		Age:  2,
	}

	mmap := map[int][]Person{
		1: []Person{p1},
		2: []Person{p2},
	}
	for i := 0; i < 4; i++ {
		list := []Person{}
		list, _ = mmap[i]
		list = append(list, Person{
			Name: "p",
			Age:  i,
		})
		mmap[i] = list
	}
	return
}

func isValid(s string) bool {
	return false
}
