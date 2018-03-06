package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sounds/display"
	"strconv"
	"strings"
)

func main() {

	var inputFile, outputFile string
	flag.StringVar(&inputFile, "input", "test_1.txt", "input file")
	flag.StringVar(&outputFile, "output", "result.txt", "result file")
	flag.Parse()

	processFile(inputFile, outputFile)
}

func processFile(inPath, outPath string) {

	var width, height, size, line, lines int
	var text, result_string string

	inFile, err := os.Open(inPath)
	if err != nil {
		log.Println("Failed to open ", inPath)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(outPath)
	if err != nil {
		log.Println("Failed to open ", outPath)
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	//skip first line, not used
	scanner.Scan()

	for scanner.Scan() {

		fields := strings.Fields(scanner.Text())

		width, err = strconv.Atoi(fields[0])
		if err != nil {
			log.Println("Failed to open ", outPath)
			return
		}

		height, err = strconv.Atoi(fields[1])
		if err != nil {
			log.Println("Failed to open ", outPath)
			return
		}

		text = scanner.Text()[len(fields[0])+len(fields[1])+2 : len(scanner.Text())]
		size, lines = display.BestFontSize(width, height, text)
		line++

		result_string = "Case #" + strconv.Itoa(line+1) + ": " + strconv.Itoa(size)
		log.Println(size, lines)
		_, err = outFile.WriteString(result_string + "\n")

	}
}
