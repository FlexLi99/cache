package cache

type Cache struct {
	storage map[string]interface{}
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]interface{}),
	}
}
