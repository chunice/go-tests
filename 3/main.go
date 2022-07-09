package main

import (
	"fmt"
	"sync"
	"time"
)

type doubleLinkedList struct {
	key, val   string
	expireTime time.Time
	prev, next *doubleLinkedList
}

type cacheMap struct {
	cap, len    int
	nodes       map[string]*doubleLinkedList
	lock        *sync.RWMutex
	first, last *doubleLinkedList
}

func NewCacheMap(capacity int) *cacheMap {
	return &cacheMap{
		cap:   capacity,
		nodes: make(map[string]*doubleLinkedList),
		lock:  new(sync.RWMutex),
	}
}

func (c *cacheMap) Set(key, val string, expired time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if node, ok := c.nodes[key]; ok {
		node.val = val
		node.expireTime = time.Now().Add(expired)
		c.moveToFirst(node)
	} else {

		if c.len == c.cap {
			delete(c.nodes, c.last.key)
			c.removeLast()
		} else {
			c.len++
		}

		node = &doubleLinkedList{
			key:        key,
			val:        val,
			expireTime: time.Now().Add(expired),
		}
		c.addToFirst(node)
		c.nodes[key] = node
	}
}

func (c *cacheMap) Get(key string) string {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if node, ok := c.nodes[key]; ok {
		if time.Now().Before(node.expireTime) {
			c.moveToFirst(node)
			return node.val
		}
	}
	return ""
}

func (c *cacheMap) moveToFirst(node *doubleLinkedList) {
	if c.first == node {
		return
	} else if c.last == node {
		c.removeLast()
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	c.addToFirst(node)
}

func (c *cacheMap) removeLast() {
	if c.last.prev != nil {
		c.last.prev.next = nil
	} else {
		c.first = nil
	}
	c.last = c.last.prev
}

func (c *cacheMap) addToFirst(node *doubleLinkedList) {
	if c.last == nil {
		c.last = node
	} else {
		c.first.prev = node
		node.next = c.first
	}
	c.first = node
}

func main() {
	c := NewCacheMap(10)
	c.Set("key1", "val1", 1*time.Hour)
	val1 := c.Get("key1")
	fmt.Printf("val1: %s \n", val1)

	c.Set("key2", "val2", 1*time.Second)
	time.Sleep(2 * time.Second)
	val2 := c.Get("key2")
	fmt.Printf("val2: %s \n", val2)
}
