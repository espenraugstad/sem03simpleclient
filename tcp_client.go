package main

import (
	"net"
	"log"
	"os"
        "github.com/espenraugstad/is105sem03/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.2:8008")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("os.Args[1] = ", os.Args[1])

	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
 	_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	encryptedResponse := string(buf[:n])
        dekrypterResponse := mycrypt.Krypter([]rune(encryptedResponse), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03) - 4)
	log.Printf("reply from proxy: %s", string(dekrypterResponse))
}
