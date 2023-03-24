package main

import (
	"fmt"

	"github.com/kovalenko-tech/gomemcache"
)

func main() {
	cache := gomemcache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
