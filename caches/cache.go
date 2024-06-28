package caches

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Remove(key string)
}
