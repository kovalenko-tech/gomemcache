# GoMemCache

GoMemCache is a simple in-memory cache implementation written in Go. It provides a key-value store that can be used to cache data in memory.

### Installation
You can install GoMemCache by running the following command:

```shell
$ go get github.com/kovalenko-tech/gomemcache
```

### Usage
To use GoMemCache in your Go project, import the package and create a new instance of the `GoMemCache` type using the `New()` function:

```go
import "github.com/your-username/gomemcache"

cache := gomemcache.New()
```

You can then use the following methods to interact with the cache:

#### \`Set(key string, value interface{})`

Sets a key-value pair in the cache.

```go
cache.Set("mykey", "myvalue")
```

#### \`Get(key string) interface{}`

Retrieves the value associated with a given key from the cache.

```go
value := cache.Get("mykey")
```

#### \`Delete(key string)`

Removes a key-value pair from the cache.

```go
cache.Delete("mykey")
```

### License
GoMemCache is licensed under the MIT License. Feel free to use, modify, and distribute this code as you see fit.
