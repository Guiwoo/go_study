package memory_db

import (
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"reflect"
	"time"
)

var memoryDB *CustomDB

type CustomDB struct {
	*cache.Cache
	*logrus.Logger
	expireTime, flushTime time.Duration
}

func (c *CustomDB) Insert(name string, a any) error {
	if reflect.ValueOf(a).Kind() != reflect.Pointer {
		return ErrValueNotPointer
	}

	c.Logger.Infof("insert data on cache name : %s , data %+v\n", name, a)
	c.Set(name, a, c.expireTime)

	return nil
}

func (c *CustomDB) Clear() {
	c.Logger.Infof("flush all data in cache")
	c.Flush()
}

func (c *CustomDB) ClearExpiredValue() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.Cache.DeleteExpired()
			c.Logger.Infof("cleared expired value")
		}
	}
}

func (c *CustomDB) Get(name string) (any, error) {
	data, ok := c.Cache.Get(name)
	if !ok {
		return nil, ErrNotFound
	}
	return data, nil
}

func GetCacheDB() *CustomDB {
	return memoryDB
}

func Init(expiration, purge time.Duration) {
	if memoryDB == nil {
		memoryDB = &CustomDB{
			Cache:      cache.New(expiration*time.Second, purge*time.Second),
			Logger:     logrus.StandardLogger(),
			expireTime: expiration * time.Second,
		}
	}

	go memoryDB.ClearExpiredValue()
	return
}
