package donglindb

//DataType:数据结构类型
type DataType=uint16

//支持的数据类型
const(
	String DataType=iota
	List
	Hash
	Set
	ZSet
)
// The operations of String.
const(
	StringSet uint16 = iota
	StringRem
	StringExpire
	StringPersist
)
// The operations of List.
const (
	ListLPush uint16 = iota
	ListRPush
	ListLPop
	ListRPop
	ListLRem
	ListLInsert
	ListLSet
	ListLTrim
	ListLClear
	ListLExpire
)

// The operations of Hash.
const (
	HashHSet uint16 = iota
	HashHDel
	HashHClear
	HashHExpire
)

// The operations of Set.
const (
	SetSAdd uint16 = iota
	SetSRem
	SetSMove
	SetSClear
	SetSExpire
)

// The operations of Sorted ZSet.
const (
	ZSetZAdd uint16 = iota
	ZSetZRem
	ZSetZClear
	ZSetZExpire
)

//创建索引
func (db *donglindb)buildStringIndex()  {
	
}
//加载索引
func (db *donglindb)loadIdxFromFiles()  {
	
}
