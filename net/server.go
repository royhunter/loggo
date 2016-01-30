package main



import (
    "fmt"
    "net"
)


func main() {
    fmt.Println("hello, server")
    udpAddr := &net.UDPAddr{
    	IP: net.IPv4(127,0,0,1),
    	Port: 3333,
    }
    udpConn, err := net.ListenUDP("udp", udpAddr)
    if err != nil {
    	fmt.Println("listen failed", err)
    	return
    }

    defer udpConn.Close()

    for {
    	data := make([]byte, 2014)

    	rx_size, _, err := udpConn.ReadFromUDP(data)
    	if err != nil {
    		fmt.Println("rx failed", err)
    		continue
    	}

    	fmt.Println("rx ok, rx size is: ", rx_size)
    	fmt.Println("rx: ", string(data))
    }

}
