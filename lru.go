package main

import (
	"errors"
	"fmt"
	"math"
)

type Node struct {
	key   string
	value string
	left  *Node
	right *Node
}
type Lru struct {
	Capacity  int
	Lrumap    map[string]*Node
	EndNode   *Node
	StartNode *Node
}

func (lru *Lru) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("lru capacity should be greater than 0")
	}
	if capacity > math.MaxInt32 {
		return errors.New(fmt.Sprintf("lru capacity should be less than %d", math.MaxInt32))
	}
	lru.Capacity = capacity
	lru.Lrumap = make(map[string]*Node, lru.Capacity)
	return nil
}
func (lru *Lru) Put(key, value string) {
	if len(lru.Lrumap) == lru.Capacity { // if capacity is full
		delete(lru.Lrumap, lru.EndNode.key)
		lru.EndNode.left.right = nil
		lru.EndNode = lru.EndNode.left
	}
	//create a new node
	newnode := &Node{key, value, nil, nil}
	lru.Lrumap[key] = newnode
	lru.MoveToFront(newnode)
	if len(lru.Lrumap) == 1 { //set the end node
	    lru.EndNode = newnode
    }
}

func (lru *Lru) MoveToFront(node *Node) {
	if node.right != nil {
		node.right.left = node.left
	}
	if node.left != nil {
		node.left.right = node.right
	}
	if lru.StartNode != nil {
	    lru.StartNode.left = node
    }
	lru.StartNode = node
}

func (lru *Lru) get(key string) string {
	node := lru.Lrumap[key]
	if node == nil {
		return ""
	}
	if node != lru.StartNode {
		lru.MoveToFront(node)
	}
	return node.value
}
