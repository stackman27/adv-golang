package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/dist-cache/cache"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool // this can be determined by concensus algorithm like raft, paxos etc
}

type Server struct {
	ServerOpts

	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error: %s", err)
	}

	log.Printf("server starting on port [%s]\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %s\n", err)
			continue
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("conn read error: %s\n", err)
			break
		}

		msg := buf[:n]
		fmt.Println(string(msg))
	}
}

func (s *Server) handleCommand(conn net.Conn, rawCmd []byte) {
	var (
		rawStr = string(rawCmd)
		parts  = strings.Split(rawStr, " ")
	)

	if len(parts) == 0 {
		// respond
		log.Println("invalid command")
		return
	}

	cmd := Command(parts[0])
	if cmd == CMDSet {
		if len(parts) != 4 {
			// respons
			log.Println("invalid set command")
			return
		}

		ttl, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Println("Invalid set command ")
			return
		}

		msg := MSGSet{
			Key:   []byte(parts[1]),
			Value: []byte(parts[2]),
			TTL:   time.Duration(ttl),
		}
		if err := s.handleSetCmd(conn, msg); err != nil {
			// respond

			return
		}

	}
}

func (s *Server) handleSetCmd(conn net.Conn, msg MSGSet) error {
	return nil
}
