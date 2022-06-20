package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}
	filelist, err := filepath.Glob(path) //
	if err != nil {
		fmt.Println("Wrong Path err : ", err, "path : ", path) //
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0

	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch)
	}

	for findInfo := range ch {
		findInfos = append(findInfos, findInfo) //
		recvCnt++
		if recvCnt == cnt {
			break
		}
	}
	return findInfos
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Wrong File err : ", err, "path : ", file) //
		ch <- findInfo                                         //
		return
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	ch <- findInfo //
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Need More Than 2 Argument")
		return
	}

	word := os.Args[1]   //target word
	files := os.Args[2:] //target file
	findInfos := []FindInfo{}

	for _, path := range files { //path에 있는 file들을 전부 search
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("-----------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("-----------------------------")
		fmt.Println()
	}
}
