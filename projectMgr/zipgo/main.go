package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
)

const BASE_DIR string = "/mnt/g/baza_mgr_large/test_file_large.docx"

func main() {
	reader, err := zip.OpenReader(BASE_DIR)
	if err != nil {
		fmt.Println("Error opening zip", err)
		return
	}
	defer reader.Close()
	buffer := make([]byte, 1024*32)
	for _, f := range reader.File {
		if f.Name != "word/document.xml" {
			continue
		}
		rdr, err := f.Open()
		if err != nil {
			fmt.Println("Error on opening word/doc", err)
			return
		}
		defer rdr.Close()
		n, err := rdr.Read(buffer)
		if err != nil {
			fmt.Println("Error on reading from buffer", err)
			return
		}
		buff := buffer[:n]
		twoDbuffer := bytes.SplitAfter(buff, []byte("<w:t>"))
		for _, tdb := range twoDbuffer[:] {
			n := bytes.Index(tdb, []byte("</w:t>"))
			if n != -1 {
				plain_text := tdb[:n]
				os.Stdout.Write(plain_text)
			}
		}
	}
}
