package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)

    record := make(map[string][]string)

    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, record, "")
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }

            countLines(f, counts, record, arg)
            f.Close()
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\t", n, line)
            fmt.Printf("Included in: %s\n", strings.Join(record[line], " "));
        }
    }
}

func countLines(f *os.File, counts map[string]int, record map[string][]string, filename string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        line := input.Text()
        counts[line]++
        if len(record[line]) == 0 {
            record[line] = append(record[line], filename)
        }

        if record[line][len(record[line])-1] !=  filename {
            record[line] = append(record[line], filename)
        }
    }
}

