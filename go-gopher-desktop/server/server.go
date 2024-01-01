package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

const (
	imageName = "imageName%d"
	imagePath = "/Users/apple/Downloads/"
)

var (
	m = map[string]string{
		"imageName1": "smy.png",
		"imageName2": "br.png",
	}
)

func main() {
	http.HandleFunc("/random", handleRandomRequest)
	fmt.Println("http server listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRandomRequest(w http.ResponseWriter, r *http.Request) {
	random := rand.Intn(2) + 1
	targetImageName := fmt.Sprintf(imageName, random)
	f, _ := os.OpenFile(imagePath+m[targetImageName], os.O_RDONLY, 0666)
	buf, _ := io.ReadAll(f)
	w.Write(buf)
}
