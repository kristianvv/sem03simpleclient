package main

import (
	"log"
	"net"
	"os"
)

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789 .,:;")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], alphabet)
		if indeks == -1 {
			kryptertMelding[i] = melding[i]
			continue
		}
		if indeks+chiffer >= len(alphabet) {
			kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]
		} else {
			kryptertMelding[i] = alphabet[indeks+chiffer]
		}
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
			break
		}
	}
	return -1
}

func main() {
	conn, err := net.Dial("tcp", "172.17.0.2:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("os.Args[1] = ", os.Args[1])

	message := []rune(os.Args[1])
	encrypted := Krypter(message, ALF_SEM03, 5)
	_, err = conn.Write([]byte(string(encrypted)))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	response := string(buf[:n])
	decrypted := Krypter([]rune(response), ALF_SEM03, -5)
	log.Printf("reply from server: %s", string(decrypted))
}
