package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	newHead := &Node{}
	tmpNewHead := newHead
	m := make(map[*Node]*Node, 0)
	tmp := head
	for tmp != nil {
		v := new(Node)
		v.Val = tmp.Val
		m[tmp] = v
		tmpNewHead.Next = v
		tmpNewHead = tmpNewHead.Next
		tmp = tmp.Next
	}
	tmp = head
	tmpNewHead = newHead.Next
	for tmp != nil {
		tmpNewHead.Random = m[tmp.Random]

		tmp = tmp.Next
		tmpNewHead = tmpNewHead.Next
	}
	return newHead.Next
}

func main() {

}
