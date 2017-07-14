// clientTCP project main.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"syscall"
	"time"
)

func main() {
	//bytes1 := []byte{0x01, 0x03, 0x00, 0x40, 0x00, 0x02, 0xC5, 0xDF}
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Text to send: ")
	//reader.ReadString('\n')
	//fmt.Fprintf(conn, text+"\n")

	//conn.Write(bytes1)

	file, err := os.Open("Mass.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	i:=0
	i<stat.Size() {
		conn.Write(bs[0])
		duration := 3 * time.Second
		time.Sleep(duration)
	}

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Message from server: " + message)
}
