package main

import (
	"log"
	"net"
	"os"

	"github.com/kristianvv/is105sem03/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.3:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("os.Args[1] = ", os.Args[1])

	encrypted := mycrypt.Encrypt([]byte(os.Args[1]))
	_, err = conn.Write(encrypted)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	decrypted := mycrypt.Decrypt(buf[:n])
	response := string(decrypted)
	log.Printf("reply from server: %s", response)
}
