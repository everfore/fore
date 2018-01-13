package main

import (
	"fmt"
	"github.com/toukii/gosseract"

	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println(err)
	}
	pprof.StartCPUProfile(f)

	client := gosseract.NewClient()
	defer client.Close()
	// client.SetImage("timg.jpg")
	client.SetLanguage("chi_sim")
	client.SetImage("ocr.png")
	// client.SetImage("text.png")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!

	pprof.StopCPUProfile()
}
