package main

func main() {
	gRPCServer := NewGRPCServer(":8000")

	
	gRPCServer.Run()
}