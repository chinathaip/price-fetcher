package main

func main() {
	svc := NewLoggerService(&priceFetcherService{})

	srv := NewJSONAPIServer(":3000", svc)
	srv.Run()
}
