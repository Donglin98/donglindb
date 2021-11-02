package storage

type(

	Meta struct {
		Key 		[]byte
		Value 		[]byte
		Extra 		[]byte	//entry额外信息
		KeySize	 	uint32
		ValueSize	uint32
		ExtraSize	uint32
	}
)
