package main

import (
	"flag"
)

func main() {

	// client := client.New("http://137.184.34.49:3000") 

	// price, err := client.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)
	
	// return 

	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running") 
	flag.Parse()

	pf := &pricefetcher{}
	svc := NewLoginService(NewMetricService(pf))
	server := NewJSONAPIServer(*listenAddr, svc)
	
	server.Run()
}