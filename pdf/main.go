package main

import (
	"fmt"
	"image/png"
	"os"

	// "github.com/jung-kurt/gofpdf"
	// "github.com/signintech/gopdf"
	pdf2 "github.com/Soreil/pdf"
	"rsc.io/pdf"
)

var fileName = "go.pdf"

func rscPDF() {

	r, err := pdf.Open(fileName)
	fmt.Println(err)

	c := r.Page(0).Content()
	for _, txt := range c.Text {
		fmt.Println(txt.S)
	}
}

func JKPDF() {
	// _ = gofpdf.PointType
}

func goPDF() {
	// gopdf.
}

func pdfT() {
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	img, err := pdf2.Decode(f)
	if err != nil {
		fmt.Println(err)
	}
	if img == nil {
		panic("decode img is nil")
	}

	imgFile, err := os.OpenFile("test.png", os.O_CREATE|os.O_WRONLY, 0644)
	defer imgFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	err = png.Encode(imgFile, img)
	// err = jpeg.Encode(imgFile, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// JKPDF()
	pdfT()
}
