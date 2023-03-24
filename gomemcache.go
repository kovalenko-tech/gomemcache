package gomemcache

type GoMemCache struct {
	d map[string]interface{}
}

func New() *GoMemCache {
	return &GoMemCache{
		d: make(map[string]interface{}),
	}
}

func (c *GoMemCache) Set(key string, value interface{}) {
	c.d[key] = value
}

func (c *GoMemCache) Get(key string) interface{} {
	return c.d[key]
}

func (c *GoMemCache) Delete(key string) {
	delete(c.d, key)
}
