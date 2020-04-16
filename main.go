package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
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
		duration := getDuration(getTimeTag(string(line)))

		time.Sleep(duration)
	}
}

// Regex syntax: https://github.com/google/re2/wiki/Syntax
func getTimestamp(timeTag string) string {
	timeRegex := regexp.MustCompile(`[0-9]{2}:[0-5][0-9]:[0-5][0-9],[0-9]{3}`)
	timestamp := timeRegex.FindString(timeTag)
	return timestamp
}

func getDuration(timeTag string) time.Duration {
	timeRegex := regexp.MustCompile(`\s[0-5][0-9],[0-9]{3}`)
	timestamp := timeRegex.FindString(timeTag)
	strDuration := strings.TrimLeft(timestamp, " ")
	seconds, _ := strconv.ParseInt(strDuration[:strings.Index(strDuration, ",")], 10, 64)
	milliseconds, _ := strconv.ParseInt(strDuration[strings.Index(strDuration, ",")+1:], 10, 64)
	return (time.Duration(milliseconds) * time.Millisecond) + (time.Duration(seconds) * time.Second)
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
