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

func getNumberMap() map[string]string {
    return map[string]string{
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }
}

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

func parseLineByChars(line string) int {
    var numbers []string
    numberMap := getNumberMap()
    for i := 0; i < len(line); i++ {
        if unicode.IsDigit(rune(line[i])) {
            numbers = append(numbers, string(line[i]))
        }
        for name, num := range numberMap {
            if strings.HasPrefix(line[i:], name) {
                numbers = append(numbers, num)
            }
        }
	}

    combinedDigits := fmt.Sprintf("%s%s", numbers[0], numbers[len(numbers)-1])
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
        number := parseLineByChars(line)
        total += number
        fmt.Printf("Number: %d\n", number)
    }
    fmt.Printf("Total: %d\n", total)
}

