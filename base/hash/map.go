package hash

type Node struct {
	key   string
	value interface{}
	next  *Node
}

type Map struct {
	size  int
	table []*Node
}

func NewHashMap() *Map {
	return &Map{
		size:  0,
		table: make([]*Node, 16),
	}
}

func (hm *Map) Add(key string, value interface{}) {
	index := hash(key)
	node := hm.table[index]

	for node != nil {
		if node.key == key {
			node.value = value
			return
		}
		node = node.next
	}

	newNode := &Node{
		key:   key,
		value: value,
		next:  hm.table[index],
	}

	hm.table[index] = newNode
	hm.size++
}

func (hm *Map) Get(key string) interface{} {
	index := hash(key)

	node := hm.table[index]
	for node != nil {
		if node.key == key {
			return node.value
		}
		node = node.next
	}

	return nil
}

func (hm *Map) Remove(key string) {
	index := hash(key)

	node := hm.table[index]
	var prev *Node
	for node != nil {
		if node.key == key {
			if prev == nil {
				hm.table[index] = node.next
			} else {
				prev.next = node.next
			}
			hm.size--
			return
		}
		prev = node
		node = node.next
	}
}

func (hm *Map) Size() int {
	return hm.size
}

func (hm *Map) IsEmpty() bool {
	return hm.size == 0
}

func (hm *Map) Clear() {
	hm.size = 0
	hm.table = make([]*Node, 16)
}

func (hm *Map) Range(f func(key string, value interface{})) {
	for _, node := range hm.table {
		for node != nil {
			f(node.key, node.value)
			node = node.next
		}
	}
}

func (hm *Map) Keys() []string {
	keys := make([]string, 0, hm.size)
	hm.Range(func(key string, value interface{}) {
		keys = append(keys, key)
	})
	return keys
}

func (hm *Map) Values() []interface{} {
	values := make([]interface{}, 0, hm.size)
	hm.Range(func(key string, value interface{}) {
		values = append(values, value)
	})
	return values
}

func (hm *Map) ContainsKey(key string) bool {
	return hm.Get(key) != nil
}

func hash(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = (h << 5) | (h >> 27)
		h += int(key[i])
	}
	return h % 16
}
