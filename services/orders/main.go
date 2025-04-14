package main

func main() {
	gRPCServer := NewGRPCServer(":8000")
	httpServer := NewHttpServer(":8080")

	go httpServer.Run()

	gRPCServer.Run()
}