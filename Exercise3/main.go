package main


import (
	"fmt"
	"net"
	"time"
)

//Finding the IP (by receiving UDP packets)
func udp_receive(port string) {
	address, resolve_error := net.ResolveUDPAddr("udp",":"+port)
	if resolve_error != nil{
		fmt.Printf("Error: %v", resolve_error)
		return
	}
	connection_p, listen_error := net.ListenUDP("udp",address)
	if listen_error != nil{
		fmt.Printf("Error: %v", listen_error)
		return
	}
	buffer := make([]byte, 1024)

	n, _, _ := connection_p.ReadFromUDP(buffer)
	fmt.Println(string(buffer[:n]))
	
	defer connection_p.Close()
}
//IP is 10.100.23.147

//Sending UDP packets
func udp_send(message, port string){
	listen_address, resolve_error := net.ResolveUDPAddr("udp4",":"+port)
	if resolve_error != nil{
		fmt.Printf("Error: %v", resolve_error)
		return
	}
	
	connection_p, listen_error := net.ListenUDP("udp4",listen_address)
	if listen_error != nil{
		fmt.Printf("Error: %v", listen_error)
		return
	}
	
	defer connection_p.Close()

	send_address, resolve_error := net.ResolveUDPAddr("udp4","10.100.23.147:"+port)
	if resolve_error != nil{
		fmt.Printf("Error: %v", resolve_error)
		return
	}

	_, write_error := connection_p.WriteToUDP([]byte(message), send_address)
	if write_error != nil{
		fmt.Printf("Error: %v", write_error)
		return
	}


}

func tcp_connect(message, port string) {
	address, resolve_error := net.ResolveTCPAddr("tcp", "10.100.23.147:"+port)
	if resolve_error != nil{
		fmt.Printf("Error: %v", resolve_error)
		return
	}

	connection_p, dial_error := net.DialTCP("tcp", nil, address)
	if dial_error != nil{
		fmt.Printf("Error: %v", dial_error)
		return
	}


	buffer := make([]byte, 1024)

	n, _ := connection_p.Read(buffer)
	fmt.Println(string(buffer[:n]))

	for {
		connection_p.Write([]byte(message+"\000"))
		
		n, _ := connection_p.Read(buffer)
		fmt.Println(string(buffer[:n]))

		time.Sleep(1*time.Second)
	}
	
}



func main(){
	/*
	udp_receive("30000")
	for {
		udp_send("Hello world!", "20013")
		udp_receive("20013")
		time.Sleep(1*time.Second)
	}
	*/

	tcp_connect("Connect to: 10.100.23.192:34933","34933")
}