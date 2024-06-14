package protected

import (
	"sync"
	"testing"
)

// TestProtected tests the basic functionality of the Protected type.
func TestProtected(t *testing.T) {
	// Initialize with an integer
	protInt := New(10)

	// Test Get method
	val := protInt.Get()
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	// Test Set method
	protInt.Set(20)
	val = protInt.Get()
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	// Test Update method
	protInt.Update(func(val int) int {
		return val * 2
	})
	val = protInt.Get()
	if val != 40 {
		t.Errorf("Expected 40, got %d", val)
	}

	// Test DoWithLock method
	protInt.DoWithLock(func(val *int) {
		*val = *val + 10
	})
	val = protInt.Get()
	if val != 50 {
		t.Errorf("Expected 50, got %d", val)
	}

	// Test DoWithRLock method
	protInt.DoWithRLock(func(val *int) {
		if *val != 50 {
			t.Errorf("Expected 50, got %d", *val)
		}
	})
}

// TestConcurrentAccess tests the concurrent access functionality of the Protected type.
func TestConcurrentAccess(t *testing.T) {
	protInt := New(0)
	var wg sync.WaitGroup
	numGoroutines := 100
	increment := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			protInt.DoWithLock(func(val *int) {
				*val++
			})
		}
	}

	// Spawn multiple goroutines to increment the value concurrently
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()

	// The final value should be numGoroutines * 1000
	expected := numGoroutines * 1000
	val := protInt.Get()
	if val != expected {
		t.Errorf("Expected %d, got %d", expected, val)
	}
}

// TestDifferentTypes tests the Protected type with different data types.
func TestDifferentTypes(t *testing.T) {
	// Test with a string
	protStr := New("hello")

	// Test Get method
	valStr := protStr.Get()
	if valStr != "hello" {
		t.Errorf("Expected 'hello', got '%s'", valStr)
	}

	// Test Set method
	protStr.Set("world")
	valStr = protStr.Get()
	if valStr != "world" {
		t.Errorf("Expected 'world', got '%s'", valStr)
	}

	// Test Update method
	protStr.Update(func(val string) string {
		return val + "!"
	})
	valStr = protStr.Get()
	if valStr != "world!" {
		t.Errorf("Expected 'world!', got '%s'", valStr)
	}

	// Test with a struct
	type myStruct struct {
		A int
		B string
	}

	protStruct := New(myStruct{A: 1, B: "initial"})

	// Test Get method
	valStruct := protStruct.Get()
	if valStruct.A != 1 || valStruct.B != "initial" {
		t.Errorf("Expected {A: 1, B: 'initial'}, got %+v", valStruct)
	}

	// Test Set method
	protStruct.Set(myStruct{A: 2, B: "updated"})
	valStruct = protStruct.Get()
	if valStruct.A != 2 || valStruct.B != "updated" {
		t.Errorf("Expected {A: 2, B: 'updated'}, got %+v", valStruct)
	}

	// Test Update method
	protStruct.Update(func(val myStruct) myStruct {
		val.A *= 2
		val.B = val.B + "!"
		return val
	})
	valStruct = protStruct.Get()
	if valStruct.A != 4 || valStruct.B != "updated!" {
		t.Errorf("Expected {A: 4, B: 'updated!'}, got %+v", valStruct)
	}
}
