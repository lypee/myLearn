package main

import "log"

type User struct {
	Name   string
	Age    int
	Height int
	Sex    int // 1=男 ;2=女
}

type userOpt func(u *User)

func WithName(name string) userOpt {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int) userOpt {
	return func(u *User) {
		u.Age = age
	}
}
func WithHeight(height int) userOpt {
	return func(u *User) {
		u.Height = height
	}
}

func Gen1(sex int, opt ...userOpt) *User {
	user1 := &User{
		Sex: sex,
	}
	for _, option := range opt {
		option(user1)
	}
	return user1
}
func main() {
	user1 := Gen1(1, WithAge(3), WithName("123"))
	log.Println(user1)
}
