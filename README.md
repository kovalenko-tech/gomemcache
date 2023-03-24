# GoMemCache

![Go Report](https://goreportcard.com/badge/github.com/informitas/cache)
![Repository Top Language](https://img.shields.io/github/languages/top/informitas/cache)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/informitas/cache)
![Github Repository Size](https://img.shields.io/github/repo-size/informitas/cache)
![Github Open Issues](https://img.shields.io/github/issues/informitas/cache)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/informitas/cache)

GoMemCache is a simple in-memory cache implementation written in Go. It provides a key-value store that can be used to cache data in memory.

### Installation
You can install GoMemCache by running the following command:

```shell
$ go get github.com/kovalenko-tech/gomemcache
```

### Usage
To use GoMemCache in your Go project, import the package and create a new instance of the `GoMemCache` type using the `New()` function:

```go
import "github.com/kovalenko-tech/gomemcache"

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
