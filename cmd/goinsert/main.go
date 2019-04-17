package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func proc(file string, lineno int, text string) error {
	_, err := os.Stat(file)
	if err != nil {
		return err
	}
	fp, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	// create new text
	c := 0
	r := bufio.NewReader(fp)
	var buf bytes.Buffer
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(line, "\n")
		if c == lineno {
			buf.WriteString(text)
			buf.WriteRune('\n')
		}
		buf.WriteString(line)
		buf.WriteRune('\n')
		c++
	}
	fp.Close()
	if c == 0 {
		buf.WriteString(text)
		buf.WriteRune('\n')
	}
	return ioutil.WriteFile(file, buf.Bytes(), 0644)
}

func main() {

	var (
		line = flag.Int("line", 0, "lineIndex")
		text = flag.String("text", "", "text")
		file = flag.String("file", "", "filePath")
	)
	flag.Parse()
	exe, _ := os.Executable()
	fmt.Println(exe)
	// read args
	lineVal := *line
	textVal := *text
	fileVal := *file
	// read target files
	if fileVal == "" {
		var buf bytes.Buffer
		for _, arg := range flag.Args() {
			_, err := os.Stat(arg)
			if err == nil {
				buf.WriteString(arg)
				buf.WriteRune(' ')
			}
		}
		fileVal = buf.String()
	}
	words := strings.Split(fileVal, " ")
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		if err := proc(word, lineVal, textVal); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
	fmt.Println(lineVal)
	fmt.Println(textVal)
}
