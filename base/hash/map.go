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

// 这个方法是实现一个哈希函数，将字符串映射为一个0到15的整数，可以用来将字符串存储在散列表中，便于快速查找。具体原理如下：
//
//首先将哈希值h初始化为0，然后遍历字符串的每一个字符（使用for循环）。
//对于每一个字符，将哈希值左移5位（h<<5），相当于将h乘以2的5次方，然后将哈希值右移27位（h>>27），相当于将h除以2的27次方，这样可以保证哈希函数分布均匀。然后将结果与字符的ASCII码相加，得到新的哈希值。
//最后我们希望将哈希值变为0到15的整数，所以使用模运算（%16）将哈希值映射到0到15的范围内。
//综合以上三个步骤，得到的就是字符串的哈希值，能够实现高效的散列表存储、查找操作。
func hash(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = (h << 5) | (h >> 27)
		h += int(key[i])
	}
	return h % 16
}
