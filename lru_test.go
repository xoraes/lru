package main

import (
    "math"
    "strconv"
    "testing"
)

func TestInitCapacitySizeLru(t *testing.T) {
    lru := &Lru{}
    err := lru.Init(0)
    if err == nil {
        t.Error("test failed")
    }
    err = lru.Init(math.MaxInt32+1)
    if err == nil {
        t.Error("test failed")
    }
}

func TestFrontValue(t *testing.T) {
    lru := &Lru{}
    lru.Init(3)
    for i := 0; i < 4; i++ {
        lru.Put(strconv.Itoa(i),strconv.Itoa(i))
    }
    lru.Put("1","1")
    v := lru.StartNode.key
    if v != "1" {
        t.Error("Failing - front value test")
    }
    lru.get("3")
    ww := lru.StartNode.key
    if ww != "3" {
        t.Error("Failing - front value test")
    }

}
func TestSingleValue(t *testing.T) {
    lru := &Lru{}
    lru.Init(10)
    lru.Put(strconv.Itoa(1),strconv.Itoa(1))
    v := lru.get("1")
    if v != "1" {
        t.Error("key/val in cache. wrong value returned")
    }
    w := lru.get("0")
    if w != "" {
        t.Error("key/value not in cache. wrong value returned")
    }
}
func TestLruFunction(t *testing.T) {
    lru := &Lru{}
    lru.Init(10)
    for i := 0; i < 100; i++ {
        lru.Put(strconv.Itoa(i),strconv.Itoa(i))
    }
   v := lru.get("99")
   if v != "99" {
       t.Error("key/val in cache. wrong value returned")
   }
    w := lru.get("1")
    if w != "" {
        t.Error("key/value not in cache. wrong value returned")
    }
}
