package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
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

		println(getText(string(line)))
		println()

		time.Sleep(time.Second)
	}
}

// Regex syntax: https://github.com/google/re2/wiki/Syntax
func getTimestamp(timeTag string) string {
	timeRegex := regexp.MustCompile(`[0-9]{2}:[0-5][0-9]:[0-5][0-9],[0-9]{3}`)
	timestamp := timeRegex.FindString(timeTag)
	return timestamp
}

func getDuration(timeTag string) string {
	timeRegex := regexp.MustCompile(`\s[0-5][0-9],[0-9]{3}`)
	timestamp := timeRegex.FindString(timeTag)
	return strings.TrimLeft(timestamp, " ")
}

func getTimeTag(line string) string {
	start := strings.Index(line, "[[")
	end := strings.Index(line, "]]")
	return line[start+2 : end]
}

func getText(line string) string {
	timeTag := getTimeTag(line)
	start := len(timeTag) + 5
	return line[start:]
}
