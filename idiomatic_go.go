package main

import (
	"fmt"
	"net"
	"sync"
)

// idiomatic golang

// const
const (
	scalar = 1.0
	vector = 2.0
)

// Panic therefore funcNum has to start with MUST
func MustParseInt(s string) {
	panic("oops")
} 

// struct 
// always call like this: 	vector := &Vector{x: 10, y: 20}
type Vector struct {
	x, y int 
}

// mutex grouping 
type mutexStruct struct {
	listenAddr string 
	isRunning bool
	
	// maps aren't concurrent safe so we protect it with mutex 
	mu sync.RWMutex
	peers map[string]net.Conn
}

// interface declaration/naming 
// - compose interface 
type Getter interface {
	Get() 
}

type Setter interface {
	Set()
}

type Storer interface {
	Getter
	Setter
}

// function to not type default return 
// this returns the default int value
func Add(x,y int) (sum int) {
	return 
}


func idomatic_main() {
	fmt.Println("hello world")
}