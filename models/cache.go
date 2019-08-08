package models

var Cache = NewCache()

type LocalCache struct {
	data map[string]interface{}
}

func (l *LocalCache) Get(key string) interface{} {
	if v, ok := l.data[key]; ok {
		return v
	}
	return ""
}

func (l *LocalCache) Put(key string, val interface{}) error {
	l.data[key] = val
	return nil
}

func (l *LocalCache) Delete(key string) error {
	delete(l.data, key)
	return nil
}

func (l *LocalCache) IsExist(key string) bool {
	if _, ok := l.data[key]; ok {
		return true
	}
	return false
}

func NewCache() *LocalCache {
	c := &LocalCache{}
	c.data = make(map[string]interface{})
	return c
}
