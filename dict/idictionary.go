package dict

type Dictionary interface {
	Clear()
	Insert(key string, val int)
	Remove(key string)
}