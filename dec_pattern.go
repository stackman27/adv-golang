package main

import (
	"fmt"
	"net/http"
)

// DECORATOR Pattern

func makeHTTPFunc(db DB, fn httpFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if err := fn(db, w, r); err != nil {
			// do something with the error
		}
	}
}

func Decmain() { 
	s := &store{}

	http.HandleFunc("/", makeHTTPFunc(s, handler)) 
	Execute(myExecuteFunc(s))
}


type DB interface {
	Store(string) error 
}

type store struct {}

func (s *store) Store(str string) error {
	fmt.Println("storing into db", str)
 
	return nil
}

func handler(db DB, w http.ResponseWriter, r *http.Request) error {
	return nil 
}

type httpFunc func(db DB, w http.ResponseWriter, r *http.Request) error

// this is coming from a third party library
type ExecuteFn func(string) 

func myExecuteFunc (db DB) ExecuteFn{
	// access to db 
	return func (s string) {
		fmt.Println("my ex functoin")
		db.Store(s) 
	}
}

 
func Execute(fn ExecuteFn) {
	fn("fooBAAR")
}

 