package storage

import (
	"github.com/roseduan/mmap-go"
	"os"
)
//db文件读写的方法
type FileRWMethod uint8
// DBFile define the data file of rosedb.
type DBFile struct {
	Id     uint32
	Path   string
	File   *os.File
	mmap   mmap.MMap
	Offset int64
	method FileRWMethod
}
