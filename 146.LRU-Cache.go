import "container/list"

type LRUCache struct {
	l *list.List
	m map[int](*list.Element)
	c int
}
type item struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	l := list.New()
	return LRUCache{
		l: l,
		m: make(map[int](*list.Element)),
		c: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	v, exist := this.m[key]
	if !exist {
		return -1
	}
	this.l.MoveToFront(v)
	return v.Value.(item).value
}

func (this *LRUCache) Put(key int, value int) {
	i := item{
		key:   key,
		value: value,
	}
	if e, exist := this.m[key]; exist {
		e.Value = i
		this.l.MoveToFront(e)
	} else {
		if this.l.Len() < this.c {
			e := this.l.PushFront(i)
			this.m[key] = e
		} else {
			e := this.l.Back()
			this.l.Remove(e)
			delete(this.m, e.Value.(item).key)
			e = this.l.PushFront(i)
			this.m[key] = e
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
