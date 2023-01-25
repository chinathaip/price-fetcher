package main

func main() {
	// c := client.New("http://localhost:3000")
	// resp, err := c.FetchPrice(context.Background(), "BTC")
	// if err != nil {
	// 	log.Fatalf("Error : %v", err)
	// }
	// fmt.Println(resp)
	// return

	svc := NewLoggerService(&priceFetcherService{})

	srv := NewJSONAPIServer(":3000", svc)
	srv.Run()
}
