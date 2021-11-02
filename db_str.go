package donglindb

import (
	"donglindb/idx"

	"errors"
	"sync"
)
//定义一些恒定的变量

var(
	ErrEmptyKey =errors.New("donglindb:the key is Empty")
	ErrKeyTooLarge = errors.New("rosedb: key exceeded the max length")
	ErrValueTooLarge = errors.New("rosedb: value exceeded the max length")
)

//sting index struct
type StrIdx struct {
	mu			*sync.RWMutex
	idxList		*idx.SkipList
}

//new string index
func NewStrdIdx() *StrIdx{
	return &StrIdx{
		mu: new(sync.RWMutex),
		idxList: idx.NewSkipList(),
	}
}
//根据key添加值，如重复则覆盖
//添加成功后，将放弃以前与密钥相关的任何生存时间。
func (db *DonginDB) Set(key []byte,value interface{}) error{
		 encKey,enValue,err := db.encode(key,value)
		 if err !=nil{
			return err
		 }
		 return db.setValue(encKey,enValue)
}

func (db *DonginDB) setValue(key []byte,value ...[]byte)( error) {
		keysize :=uint32(len(key))
		if keysize ==0{
			return ErrEmptyKey
		}
		config := db.config
		if keysize >config.MaxKeySize{
			return ErrKeyTooLarge
		}
		for _,v :=range value {
			if uint32(len(v)) > config.MaxValueSize {
				return ErrValueTooLarge
			}
		}
		return nil
}
