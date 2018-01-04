package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	// var png []byte
	// png, err := qrcode.Encode("http://192.168.1.102:8080", qrcode.Medium, 256)

	err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("output qr.png")
	}
}
