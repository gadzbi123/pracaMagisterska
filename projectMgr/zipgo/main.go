package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

const BASE_DIR string = "/run/media/gadzbi/GryIFilmy/baza_mgr_large/test_file_large.docx"

/*
Ok lets think, should I move archives to new folder, which one?
This one will have more sense to mesure perfomance but it requires precreation of folders

Or should i make archive unzips during algo run?
This might make it slower for algo (who care???),
use more mem (can fix with static buffer),
easier to do becuse i have code below,
harder to mesure archives??? I can check sizes of files anyway
*/
func main() {
	reader, err := zip.OpenReader(BASE_DIR)
	if err != nil {
		fmt.Println("Error opening zip", err)
		return
	}
	defer reader.Close()
	buffer := make([]byte, 1024*32)

	// this goes through all the files, ignores dirs so i don't need to care about them
	for _, f := range reader.File {
		// fmt.Printf("Name: %v, size: %v, dir?: %v\n", f.Name, f.CompressedSize64, f.FileInfo().IsDir())
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
		twoDbuffer := bytes.SplitAfterN(buff, []byte("<w:t>"), 1)
		for _, tdb := range twoDbuffer[:] {
			n := bytes.Index(tdb, []byte("</w:t>"))
			if n != -1 {
				// plain_text := tdb[:n]
				// os.Stdout.Write(plain_text)
			}
		}
	}
}
