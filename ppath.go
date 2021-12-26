package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var colorMap = map[string]string{
	"RESET":  "\033[0m",
	"BOLD":   "\033[1m",
	"GRAY":   "\033[38;5;243m",
	"RED":    "\033[38;5;202m",
	"GREEN":  "\033[38;5;35m",
	"CYAN":   "\033[38;5;75m",
	"ORANGE": "\033[38;5;202m",
	"BROWN":  "\033[38;5;94m",
}

func contains(array []string, target string) bool {
	for _, element := range array {
		if target == element {
			return true
		}
	}
	return false
}

func unique(array []string) []string {
	seen := []string{}
	for _, element := range array {
		if contains(seen, element) {
			continue
		} else {
			seen = append(seen, element)
		}
	}
	return seen
}

func printlnColor(color string, element string) {
	fmt.Println(colorMap["BOLD"] + colorMap[color] + element + colorMap["RESET"])
}

func colorizePath(element string) {
	defaultPaths := []string{"/bin", "/sbin", "/usr/bin", "/usr/sbin", "/usr/local/bin", "/usr/local/sbin"}

	if contains(defaultPaths, element) {
		printlnColor("GRAY", element)
	} else if strings.Contains(element, "/go/") || strings.HasSuffix(element, "/go") {
		printlnColor("CYAN", element)
	} else if strings.Contains(element, "venv") || strings.Contains(element, "conda") || strings.Contains(element, "python") || strings.Contains(element, "Python") || strings.Contains(element, "py") {
		printlnColor("GREEN", element)
	} else if strings.Contains(element, "homebrew") || strings.Contains(element, "linuxbrew") {
		printlnColor("BROWN", element)
	} else if strings.Contains(element, "/.cargo/") {
		printlnColor("ORANGE", element)
	} else {
		printlnColor("BOLD", element)
	}
}

func main() {
	allPtr := flag.Bool("all", false, "Include duplicate paths from $PATH variable. (default false)")
	uncoloredPtr := flag.Bool("uncolored", false, "Uncolor the output. (default false)")
	flag.Parse()
	allArg := *allPtr
	uncoloredArg := *uncoloredPtr

	stringPaths := os.Getenv("PATH")
	paths := strings.Split(stringPaths, ":")

	if !allArg {
		paths = unique(paths)
	}

	if uncoloredArg {
		for _, element := range paths {
			fmt.Println(element)
		}
	} else {
		for _, element := range paths {
			colorizePath(element)
		}
	}

}
