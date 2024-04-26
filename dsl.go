package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
)

func dsl(pattern string, showMatching bool) {
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
        // If the line matches the pattern and showMatching is true, print it
        // If the line does not match the pattern and showMatching is false, print it
        if regex.MatchString(line) == showMatching {
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
    // Define command-line flags
    showMatching := flag.Bool("s", false, "Show only lines that match the pattern")
    flag.Parse()

    // Extract the pattern from command-line arguments
    args := flag.Args()
    if len(args) < 1 {
        fmt.Println("Usage: dsl [-s] <pattern>")
        os.Exit(1)
    }
    pattern := args[0]

    // Call dsl function with the pattern and showMatching flag
    dsl(pattern, *showMatching)
}
