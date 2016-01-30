package main


import (
    "fmt"
    "net"
)

func main() {
    fmt.Println("hello, client")

    udpAddr := &net.UDPAddr{
    	IP: net.IPv4(127,0,0,1),
    	Port: 3333,
    }

    udpConn, err := net.DialUDP("udp", nil, udpAddr)
    if err != nil {
    	fmt.Println("connect failed")
    	return
    }
    fmt.Println("connect ok")

    defer udpConn.Close()

    data := []byte("hello")

    tx_size, err := udpConn.Write(data)
    if err != nil {
    	fmt.Println("tx error:", err)
    	return
    }

    fmt.Println("send ok, tx size is ", tx_size)
}
