package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const InputFilePath = "./input.txt"

func readLinesOfFile(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func parseGame(line string) {
    setString := strings.Split(line, ":")
    sets := strings.Split(setString[1], ";")

    for _, set := range sets {
        cubes
    }
}

func main() {
    lines, err := readLinesOfFile(InputFilePath)
    if err != nil {
        panic(err)
    }

    total := 0
    for _, line := range lines {

    }
    fmt.Printf("Output: %d\n", total)
}

