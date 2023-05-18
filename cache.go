package cache



type Cache struct {
	Data map[string]interface{}
}

func NewCache() Cache {
	return Cache{
	make (map[string]interface{}),
	}
}

func (c *Cache) Get(key string)	interface{} {
	return c.Data[key]
}

func (c *Cache) Set(key string, value interface{}) {
	c.Data[key] = value
}	 	

func (c *Cache) Delete(key string) {
	delete(c.Data, key)
}	
	
