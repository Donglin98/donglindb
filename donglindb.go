package donglindb

import (
	"donglindb/storage"
	"donglindb/utils"
	"sync"

)

//东临数据库结构，表示一个数据库的实例

type (
	DonginDB struct {
		activeFile *sync.Map
		archFiles  ArchivedFiles
		strIndex        *StrIdx       // String indexes(a skip list).
		//	listIndex       *ListIdx      // List indexes.
		//	hashIndex       *HashIdx      // Hash indexes.
		//	setIndex        *SetIdx       // Set indexes.
		//	zsetIndex       *ZsetIdx      // Sorted set indexes.
		config          Config        // Config info of rosedb.
		mu              sync.RWMutex  // mutex.
		expires         Expires       // Expired directory.
		isMerging       bool          // Indicates whether the db is merging, see StartMerge.
		isSingleMerging bool          // Indicates whether the db is in single merging, see SingleMerge.
	//	lockMgr         *LockMgr      // lockMgr controls isolation of read and write.
	//	txnMeta         *TxnMeta      // Txn meta info used in transaction.
		closed          uint32
		mergeChn        chan struct{} // mergeChn used for sending stop signal to merge func.
	}
	//ArchivedFiles 归档文件
	ArchivedFiles map[DataType]map[uint32]*storage.DBFile
	//Expires saves the expire info of different keys.
	Expires map[DataType]map[string]int64
)





//字节编译
func (db *DonginDB) encode(key,value interface{}) (encKey,enValue []byte,err error) {
	encKey,err = utils.EncodeKey(key)
	if err !=nil{
		return
	}
	if enValue, err = utils.EncodeValue(value); err != nil {
		return
	}
	return
}
