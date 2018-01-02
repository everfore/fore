package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func main() {
	p, err := oto.NewPlayer(44100, 2, 2, 8192)
	defer p.Close()
	fmt.Println(err)

	f, _ := os.Open("sun.mp3")
	d, _ := mp3.Decode(f)
	defer d.Close()

	io.Copy(p, d)

}
