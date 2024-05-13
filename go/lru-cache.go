package leetcodesolutions

type LRUCache struct {
	capacity int
	length   int
	cache    map[int]*Node
	alist    *DoubleLinkedList
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

type DoubleLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		alist:    &DoubleLinkedList{},
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.alist.MoveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.alist.MoveToFront(node)
		return
	}

	node := &Node{key: key, value: value}
	if this.length >= this.capacity {
		lruNode := this.alist.RemoveTail()
		delete(this.cache, lruNode.key)
		this.length--
	}

	this.alist.PushToFront(node)
	this.cache[key] = node
	this.length++
}

func (dll *DoubleLinkedList) PushToFront(node *Node) {
	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		node.next = dll.head
		dll.head.prev = node
		dll.head = node
	}
	dll.length++
}

func (dll *DoubleLinkedList) RemoveTail() *Node {
	if dll.tail == nil {
		return nil
	}
	tailNode := dll.tail
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = tailNode.prev
		dll.tail.next = nil
	}
	dll.length--
	return tailNode
}

func (dll *DoubleLinkedList) MoveToFront(node *Node) {
	if node == dll.head {
		return
	}
	if node == dll.tail {
		dll.tail = node.prev
		dll.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.prev = nil
	node.next = dll.head
	dll.head.prev = node
	dll.head = node
}
