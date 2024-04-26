package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func dsl(pattern string) {
    if pattern == "" {
        fmt.Println("Usage: dsl <pattern>")
        os.Exit(1)
    }

    // Compile the regular expression pattern
    regex, err := regexp.Compile(pattern)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    // Read from standard input line by line
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        // If the line does not match the pattern, print it
        if !regex.MatchString(line) {
            fmt.Println(line)
        }
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
        os.Exit(1)
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: dsl <pattern>")
        os.Exit(1)
    }
    dsl(os.Args[1])
}

