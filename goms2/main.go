package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running") 
	flag.Parse()

	pf := &pricefetcher{}
	svc := NewLoginService(NewMetricService(pf))
	server := NewJSONAPIServer(*listenAddr, svc)
	
	server.Run()
}