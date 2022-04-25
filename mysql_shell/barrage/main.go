package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	chunkSize = 1 << (16)
	maxNum = 1000
)

var (
	basePath = "/Users/lipeiyong/go/src/lpynnng/%d.txt"
	path     = "/Users/lipeiyong/go/src/lpynnng/mysql/message.log"
	resPath  = "/Users/lipeiyong/go/src/lpynnng/mysql/result.log"
	mmap     map[string]int
	nodeList NodeList
)

type Node struct {
	Msg   string
	Times int
}

func main() {
	mmap = make(map[string]int, 0)
	nodeList = make([]Node, 0)
	for i := 1; i <= 232; i++ {
		filePath := fmt.Sprintf(basePath, i)
		ReadFile(filePath, handle)
	}
	localMap := mmap
	handleMap(localMap)
	heap.Init(&nodeList)
	localQueue := make(NodeList , maxNum)
	for i := 0 ; i < maxNum ; i++{
		localQueue[i] = nodeList[i]
	}
	WriteFile(localQueue)

	return
}

func WriteFile(nodeList NodeList) {
	var f *os.File
	if _, err := os.Stat(resPath); os.IsNotExist(err) {
		f, _ = os.Create(resPath)
	} else {
		f, _ = os.OpenFile(resPath ,os.O_APPEND | os.O_WRONLY, os.ModeAppend)

	}
	defer f.Close()
	cnt := 0
	log.Println("总行数: " , nodeList.Len())
	for {
		if nodeList.Len() > 0 && cnt < maxNum{
			node := nodeList[cnt]
			cnt++
			str := node.Msg + "   " + strconv.Itoa(node.Times)+ " \n"
			_, err := f.WriteString(str)
			if err != nil {
				log.Println("\n err!!")
				log.Println(err.Error())
				return
			}
		}else{
			break
		}
	}
	log.Println("写入行数: " ,cnt)
	f.Close()
	return
}
func ReadFile(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, _, err := buf.ReadLine()
		str := string(line)
		str = strings.TrimSpace(str)
		handle(str)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}
func handle(str string) {
	if _, exist := mmap[str]; exist {
		fmt.Println("exist")
		mmap[str]++
	} else {
		fmt.Println("no-exist")
		mmap[str] = 1
	}
}

func test() {
	student := &NodeList{{"Amy", 21}, {"Dav", 15}, {"Spo", 22}, {"Reb", 11}}
	heap.Init(student)
	one := Node{"hund", 9}
	heap.Push(student, one)
	for student.Len() > 0 {
		fmt.Printf("%v\n", heap.Pop(student))
	}
}

func handleMap(localMap map[string]int) {
	for k, v := range localMap {
		nodeList = append(nodeList, Node{
			Msg:   k,
			Times: v,
		})
	}

	return
}

func splitFile() {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	num := math.Ceil(float64(fileInfo.Size()) / chunkSize)
	fi, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	b := make([]byte, chunkSize)
	var i int64 = 1
	for ; i <= int64(num); i++ {
		fi.Seek((i-1)*chunkSize, 0)
		if len(b) > int(fileInfo.Size()-(i-1)*chunkSize) {
			b = make([]byte, fileInfo.Size()-(i-1)*chunkSize)
		}
		fi.Read(b)

		f, err := os.OpenFile("./"+strconv.Itoa(int(i))+".txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		f.Write(b)
		f.Close()
	}
	fi.Close()
}

type NodeList []Node

func (t *NodeList) Len() int {
	return len(*t) //
}

func (t *NodeList) Less(i, j int) bool {
	return (*t)[i].Times > (*t)[j].Times
}

func (t *NodeList) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *NodeList) Push(x interface{}) {
	*t = append(*t, x.(Node))
}

func (t *NodeList) Pop() interface{} {
	n := len(*t)
	x := (*t)[n-1]
	*t = (*t)[:n-1]
	return x
}
