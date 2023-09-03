package main

import "fmt"

// Efficient struct pattern (Functional Options)
// learn more: https://www.youtube.com/watch?v=MDy7JQN5MN4

type OptFunc func (*Opts) 

type Opts struct {
	maxConn int 
	id string 
	tls bool 
}

func defaultOpts() Opts {
	return Opts {
		maxConn: 10,
		id: "default",
		tls: false,
	}
}

func withId(s string) OptFunc {
	return func(opts *Opts) {
		opts.id = s
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(n int) OptFunc{
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server{
	o := defaultOpts()

	// if opts is not provided we use the defaultOpts
	// if opts are provided we can modify option pointers
	for _, fn := range opts {
		fn(&o)
	}
	 
	return &Server{
		Opts: o,
	}
}


func FuncMain() {
	// s := newServer() Default Options
	s := newServer(withTLS, withMaxConn(99), withId("sishir"))
	fmt.Println(s)
}