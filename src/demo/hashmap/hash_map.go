package hashmap

const TableSize = 16

type HashMap struct {
	Table [TableSize]*node
}

type node struct {
	key      string
	value    string
	hashCode int
	next     *node
}

func NewHashMap() HashMap {
	var hashMap HashMap
	hashMap.initTable()
	return hashMap
}

func (hashMap *HashMap) initTable() {
	for i := range hashMap.Table {
		hashMap.Table[i] = &node{hashCode: i}
	}
}

func (hashMap HashMap) genHashCode(key string) int {
	keyLen := len(key)
	if keyLen == 0 {
		return 0
	}
	var hashCode = 0
	var lastIndex = keyLen - 1
	for i := range key {
		if i == lastIndex {
			hashCode += int(key[i])
			break
		}
		hashCode += (hashCode + int(key[i])) * 31
	}
	return hashCode
}

func (hashMap HashMap) indexTable(hashCode int) int {
	return hashCode % TableSize
}

func (hashMap HashMap) indexNode(hashCode int) int {
	return hashCode / TableSize
}

func (hashMap HashMap) Put(key, value string) string {
	var hashCode = hashMap.genHashCode(key)
	var thisNode = node{key: key, value: value, hashCode: hashCode}

	var tableIndex = hashMap.indexTable(hashCode)
	var nodeIndex = hashMap.indexNode(hashCode)

	var headNode = hashMap.Table[tableIndex]

	if (*headNode).key == "" {
		*headNode = thisNode
		return ""
	}

	var lastNode = headNode
	var nextNode = (*headNode).next

	for nextNode != nil && (hashMap.indexNode((*nextNode).hashCode) < nodeIndex) {
		lastNode = nextNode
		nextNode = (*nextNode).next
	}
	if (*lastNode).hashCode == thisNode.hashCode {
		var oldValue = lastNode.value
		lastNode.value = thisNode.value
		return oldValue
	}
	if lastNode.hashCode < thisNode.hashCode {
		lastNode.next = &thisNode
	}
	if nextNode != nil {
		thisNode.next = nextNode
	}
	return ""
}

func (hashMap HashMap) Get(key string) string {
	var hashCode = hashMap.genHashCode(key)
	var tableIndex = hashMap.indexTable(hashCode)

	var thisNode = hashMap.Table[tableIndex]

	if (*thisNode).key == key {
		return (*thisNode).value
	}

	for (*thisNode).next != nil {
		thisNode = (*thisNode).next
		if key == (*thisNode).key {
			return (*thisNode).value
		}
	}
	return ""
}
