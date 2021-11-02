package donglindb

import (
	"donglindb/storage"
	"time"
)
// DataIndexMode the data index mode.
type DataIndexMode int

type Config struct {
	Addr    string `json:"addr" toml:"addr"`         // server address
	DirPath string `json:"dir_path" toml:"dir_path"` // rosedb dir path of db file
	// Deprecated: don`t edit the option, it will be removed in future release.
	BlockSize    int64                `json:"block_size" toml:"block_size"` // each db file size
	RwMethod     storage.FileRWMethod `json:"rw_method" toml:"rw_method"`   // db file read and write method
	IdxMode      DataIndexMode        `json:"idx_mode" toml:"idx_mode"`     // data index mode
	MaxKeySize   uint32               `json:"max_key_size" toml:"max_key_size"`
	MaxValueSize uint32               `json:"max_value_size" toml:"max_value_size"`

	// Sync is whether to sync writes from the OS buffer cache through to actual disk.
	// If false, and the machine crashes, then some recent writes may be lost.
	//
	// Note that if it is just the process that crashes (and the machine does not) then no writes will be lost.
	//
	// The default value is false.
	Sync bool `json:"sync" toml:"sync"`

	MergeThreshold int `json:"merge_threshold" toml:"merge_threshold"` // threshold to reclaim disk

	MergeCheckInterval time.Duration `json:"merge_check_interval"`
}
