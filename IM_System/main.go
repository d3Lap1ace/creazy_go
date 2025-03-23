package main

func main() {
	server := NewServer("0.0.0.0", 8081)
	server.Start()
}
