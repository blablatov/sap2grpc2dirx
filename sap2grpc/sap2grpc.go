// Тестовый http-сервер модуль. Выполняет обработку http-запроса, парсит принятые данные.
// передает эти данные http-post модулю, для выполнения запроса к Directum RX
package main

import (
	"drxclient"
	"fmt"

	"log"
	"net/http"
	"strings"
	"time"
)

type SapData struct {
	Id          string `json:"Id"`
	Status      string `json:"Status"`
	Name        string `json:"Name"`
	InterfaceId string `json:"InterfaceId"`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// this handler is returning component path of URL.
func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// parsing of string. Обработка URL-строки после "/" символа
	i := strings.Index(r.URL.Path, ":") // get index of symbol ":"
	sd := &SapData{}
	sd.Id = r.URL.Path[:i]       // get slice of string before symbol ":"
	sd.Status = r.URL.Path[i+1:] // get slice of string after symbol ":"
	fmt.Println("Id:", strings.TrimPrefix(sd.Id, "/"), "\nStatus:", sd.Status)

	cc := make(chan string) // канал функции gRPC-клиента
	go drxclient.PostGrpc(strings.TrimPrefix(sd.Id, "/"), sd.Status, cc)
	// Получение данных из канала cc горутины
	log.Println("\n\nРезультат запроса: Тип созданного документа: ", <-cc,
		"\nРезультат запроса: Название созданного документа: ", <-cc)
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs время выполнения запроса\n", secs)
}
