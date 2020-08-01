package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_err(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("invaders.h")
	check_err(err)
	defer file.Close()

	fileinfo, err := file.Stat()
	check_err(err)

	filesize := fileinfo.Size()
	bytes := make([]byte, filesize)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	Disassemble(bytes)

	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%#02x ", bytes[i])
	}
}
