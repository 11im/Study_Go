package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		filelist, err := GetFileList(path)

		if err != nil {
			fmt.Println("Wrong Path! :", err, "path :", path)
			return
		}

		fmt.Println("Target File List ")
		for _, name := range filelist {
			fmt.Printf("%s ", name)
		}
	}
}

func PrintFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Cannot Find File", filename)
		return

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

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

	filelist, err := GetFileList(path)

	if err != nil {
		fmt.Println("Wrong File Path : ", err, "Path :", path)
		return findInfos
	}

	for _, filename := range filelist {
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}
	return findInfos
}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot Find File ", filename)
		return findInfo
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
	return findInfo
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
