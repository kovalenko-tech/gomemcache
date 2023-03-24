package gomemcache

import "testing"

func TestGoMemCache_SetGetDelete(t *testing.T) {
	cache := New()

	// Test Set() and Get()
	cache.Set("mykey", "myvalue")
	if value := cache.Get("mykey"); value != "myvalue" {
		t.Errorf("expected %q but got %q", "myvalue", value)
	}

	// Test Delete() and Get()
	cache.Delete("mykey")
	if value := cache.Get("mykey"); value != nil {
		t.Errorf("expected nil but got %q", value)
	}
}
