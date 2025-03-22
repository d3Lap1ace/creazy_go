package main

import (
	"fmt"
	"net"
)

type Server struct{
	Ip string
	Port int
}

func NewServer(ip string,port int ) *Server{
	server := &Server{
		Ip : ip,
		Port : port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn){
	// current chain of business
	fmt.Println("success connect...")
}

func (this *Server) Start(){
	// socket listen
	listener,err := net.Listen("tcp",fmt.Sprintf("%s:%d",this.Ip,this.Port))

	if err != nil{
		fmt.Println("net.Listen err:",err)
		return
	}
	// close listen socket
	defer listener.Close()

	for{
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("net.acceot err:",err)
			return
		}


		go this.Handler(conn)
	}
}