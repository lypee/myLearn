package main

import (
	"github.com/lypee/snowFlake/server/zkServer"
	"log"

	"github.com/lypee/snowFlake"
)

var (
	sf *snowFlake.SfWorker
)

func init() {
	errCh := make(chan error, 3)
	var err error
	sf, err = snowFlake.NewSfWorker(errCh ,zkServer.WithServers([]string{"43.138.36.75:2181"}))
	if err != nil{
		log.Println(err.Error())
		return
	}

}

func main() {
	test()
}
func test() {
	id, _ := sf.NextID()
	log.Println(id)
}
