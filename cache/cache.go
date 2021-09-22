package cache

type Cache interface {
	Get(key string) (result Data, err error)
	Set(key string, object Data) (err error)
}
