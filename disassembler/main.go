package main

import (
	"bufio"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("../testdata/invaders")
	checkErr(err)
	defer file.Close()

	fileinfo, err := file.Stat()
	checkErr(err)

	filesize := fileinfo.Size()
	bytes := make([]byte, filesize)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	Disassemble(bytes)
}
