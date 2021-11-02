package idx

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

//SkipList是skip list的实现，skip list是一种高效的数据结构，可以替代平衡二叉搜索树。
//它可以以O（logN）时间复杂度平均值插入、删除和查询。
//有关跳过列表的具体说明，请参阅维基百科：https://en.wikipedia.org/wiki/Skip_list.

const (
	//skl指标的最高水平可根据实际情况进行调整。
	maxLevel int =18
	probability float64 = 1 / math.E
)
//通过迭代数据当结果为false跳出循环
type handleEle func(e *Element) bool

type(
	// Node the skip list node.
	Node struct {
		next []*Element
	}

	// Element element is the data stored.
	Element struct {
		Node
		key   []byte
		value interface{}
	}
	SkipList struct {
		Node
		maxLevel 	int
		Len			int
		randSource     rand.Source
		probability float64
		probTable   []float64
		prevNodesCache	[]*Node
	}
)

//创建一个新的跳表
func NewSkipList() *SkipList{
		return &SkipList{
			Node:			Node{next: make([]*Element,maxLevel)},
			prevNodesCache:	make([]*Node,maxLevel),
			maxLevel: 		maxLevel,
			randSource:		rand.New(rand.NewSource(time.Now().UnixNano())), //生成随机数
			probability: 	probability,
			probTable:		probabilityTable(probability,maxLevel),
		}
}
//获取下一个跳表元素信息的key
func (e *Element)Key() []byte {
	return e.key
}
//获取下一个跳表元素信息的value
func (e *Element)Value() interface{} {
	return e.value
}
//设置Value值
func (e *Element) SetValue(val interface{}) {
	e.value = val
}
//从原链表进行一个跳跃形成一个个层级索引的链表
func (t *SkipList) Front() *Element{
	return	t.next[0]
}
//跳表的一级索引是原始数据，按顺序排列。
//根据Next方法可以得到一个所有数据串联的链表。
func (e *Element) Next() *Element {
	return e.next[0]
}
//GET 根据key获取值，如果未找到则返回nil
func (t *SkipList) Get(key []byte) *Element {
	var pre = &t.Node
	var next *Element
	for i := t.maxLevel-1; i >=0 ; i-- {
		next =pre.next[i]
		for next!=nil && bytes.Compare(key,next.key)>0{
			pre = &next.Node
			next = pre.next[i]
		}
	}
	if next !=nil && bytes.Compare(next.key,key)<=0{
			return next
	}
	return nil
}
//PUT 向跳表插入值，如果key重复则覆盖
func (t *SkipList)Put(key []byte,value interface{}) *Element {
	var element *Element
	perv :=t.backNodes(key)
	if element = perv[0].next[0];element !=nil && bytes.Compare(element.key,key) <=0{
		element.value =value
		return element
	}
	element =&Element{
		Node{
			next: make([]*Element, t.randomLevel()),
		},
		key,
		value,
	}
	for i := range element.next {
		element.next[i] = perv[i].next[i]
		perv[i].next[i] = element
	}
	t.Len++
	return element
}
//Check key 是否存在
func (t *SkipList) Exist(key []byte) bool {
	return t.Get(key)!=nil
}
//Delete element by the key
func (t *SkipList) Delete(key []byte) *Element {
	perv :=t.backNodes(key)
	element := perv[0].next[0]
	if element !=nil && bytes.Compare(element.key,key)<=0{
		for k ,v :=range element.next{
			perv[k].next[k] =v
		}
		t.Len--
		return element
	}
	return nil
}

// Foreach iterate all elements in the skip list.
// 通过迭代跳表所有数据
func (t *SkipList) Foreach(ele handleEle)  {
	for p := t.Front(); p != nil; p = p.Next(){
		if ok := ele(p); !ok {
			break
		}
	}
}

// FindPrefix find the first element that matches the prefix.
func (t *SkipList) FindPrefix(prefix []byte) *Element {
	var prev = &t.Node
	var next *Element

	for i := t.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]

		for next != nil && bytes.Compare(prefix, next.key) > 0 {
			prev = &next.Node
			next = next.next[i]
		}
	}

	if next == nil {
		next = t.Front()
	}

	return next
}
//返回前一个Node节点
func (t *SkipList) backNodes(key []byte) []*Node{
	var perv =&t.Node
	var next *Element
	pervs :=t.prevNodesCache
	for i := t.maxLevel-1; i >0 ; i-- {
		next = perv.next[i]
		if next !=nil && bytes.Compare(key,next.key)>0{
			perv =&next.Node
			next =next.next[i]
		}
		pervs[i] =perv
	}
	return pervs
}
//生成随机索引级别
func (t *SkipList) randomLevel() (level int) {
	r := float64(t.randSource.Int63()) / (1 << 63)

	level = 1
	for level < t.maxLevel && r < t.probTable[level] {
		level++
	}
	return
}
func probabilityTable(probability float64, maxLevel int) (table []float64) {
	for i := 1; i <= maxLevel; i++ {
		prob := math.Pow(probability, float64(i-1))
		table = append(table, prob)
	}
	return table
}