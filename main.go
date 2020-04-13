package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	script, err := ioutil.ReadFile("./script.txt")
	if err != nil {
		log.Printf("error loading the script %v", err)
	}

	reader := bytes.NewReader(script)
	lines := bufio.NewReader(reader)

	for {
		line, _, err := lines.ReadLine()
		if err == io.EOF {
			break
		}

		println(string(line))
		println()

		time.Sleep(2 * time.Second)
	}
}
