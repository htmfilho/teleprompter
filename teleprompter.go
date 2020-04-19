package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

var scriptFile = flag.String("script", "./example.txt", "Path to the script file")

func main() {
	flag.Parse()

	script, err := ioutil.ReadFile(*scriptFile)
	if err != nil {
		log.Printf("error loading the script %v", err)
	}

	reader := bytes.NewReader(script)
	lines := bufio.NewReader(reader)
	timer := time.Duration(0)
	lineDuration := time.Duration(0)
	for {
		line, _, err := lines.ReadLine()
		if err == io.EOF {
			break
		}

		currentTimer := getTimestamp(getTimeTag(string(line)))
		duration := currentTimer - timer - lineDuration
		timer = currentTimer

		time.Sleep(duration)

		log.Println(getText(string(line)))
		fmt.Println()

		lineDuration = getDuration(getTimeTag(string(line)))

		time.Sleep(lineDuration)
	}
}

// Regex syntax: https://github.com/google/re2/wiki/Syntax
func getTimestamp(timeTag string) time.Duration {
	timeRegex := regexp.MustCompile(`[0-9]{2}h[0-5][0-9]m[0-5][0-9]s[0-9]{3}ms`)
	timestamp := timeRegex.FindString(timeTag)
	duration, _ := time.ParseDuration(timestamp)
	return duration
}

func getDuration(timeTag string) time.Duration {
	timeRegex := regexp.MustCompile(`\s[0-5][0-9]s[0-9]{3}ms`)
	timestamp := timeRegex.FindString(timeTag)
	strDuration := strings.TrimLeft(timestamp, " ")
	duration, _ := time.ParseDuration(strDuration)
	return duration
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
