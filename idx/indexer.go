package idx

import "donglindb/storage"

//Indexer the data index info, stored in skip list.
//数据索引信息，存储到跳表中
type Indexer struct {
	Meta	*storage.Meta // metadata info.
	FileId	uint32		 //入口数据查询起始位置。
	offset 	int64		//输入数据查询的起始位置
}
