package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var client http.Client

func init() {
	client = http.Client{
		Timeout: 5 * time.Second,
	}
}

func main() {
	file, err := os.Open("urls")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	contador := 1
	diretorio := "../casaCampo/"
	for scanner.Scan() {
		destino := diretorio + fmt.Sprintf("image%d.jpg", contador)
		contador += 1
		downloadImage(scanner.Text(), destino)
	}
}

func downloadImage(urlOrigem string, destino string) {

	response, err := client.Get(urlOrigem)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(destino)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fmt.Sprintf("%s Sucesso!!", destino))
}
