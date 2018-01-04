package main

import (
	"fmt"
	"github.com/otiai10/gosseract"

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
	client.SetImage("text.png")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!

	pprof.StopCPUProfile()
}
