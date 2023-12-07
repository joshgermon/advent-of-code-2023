package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func parseLine(line string) int {
    re := regexp.MustCompile("\\d")
    matches := re.FindAllString(line, -1)

    var digits []string
    for _, match := range matches {
        digits = append(digits, match)
    }
    combinedDigits := fmt.Sprintf("%s%s", digits[0], digits[len(digits)-1])
    number, _ := strconv.Atoi(combinedDigits)
    return number
}

func main() {
    lines, err := readLinesOfFile(InputFilePath)
    if err != nil {
        panic(err)
    }
    total := 0
    for _, line := range lines {
        number := parseLine(line)
        total += number
        fmt.Printf("Number: %d\n", number)
    }
    fmt.Printf("Total: %d\n", total)
}

