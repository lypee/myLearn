package priorityQueue

import "container/heap"


type ListNode struct{
	Val int
	Next *ListNode
}
type Item struct {
	value    *ListNode
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index -= 1
	*pq = old[0 : n-1]
	return item
}
func mergeKLists(lists []*ListNode) *ListNode {
	// 处理边界
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	truLength := 0
	for _, v := range lists {
		if v != nil {
			truLength ++
		}
	}
	if truLength == 0 {
		return nil
	}

	dummy := &ListNode{}
	current := dummy
	// 初始化队列
	pq := make(PriorityQueue, truLength)

	index := 0
	for _, v := range lists {
		if v != nil {
			pq[index] = &Item{
				value: v,
				priority: v.Val,
				index: index,
			}
			index ++
		}
	}
	// 堆排序
	heap.Init(&pq)

	for pq.Len() > 0 {
		// 出队
		item := heap.Pop(&pq).(*Item)
		curNode := item.value
		current.Next = curNode
		current = current.Next
		if curNode.Next != nil {
			newItem := &Item{
				value: curNode.Next,
				priority: curNode.Next.Val,
			}
			// 入队
			heap.Push(&pq, newItem)
		}
	}

	return dummy.Next

}