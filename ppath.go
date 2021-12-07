package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var colors3bit = map[string]string{
	"RESET":  "\033[0m",
	"BOLD":   "\033[1m",
	"GRAY":   "\033[30m",
	"RED":    "\033[31m",
	"GREEN":  "\033[32m",
	"CYAN":   "\033[36m",
	"ORANGE": "\033[31m",
	"BROWN":  "\033[31m",
}

var colors8bit = map[string]string{
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

func printlnColor(colorMap map[string]string, color string, element string) {
	fmt.Println(colorMap["BOLD"] + colorMap[color] + element + colorMap["RESET"])
}

func colorizePath(colorMap map[string]string, element string) {
	defaultPaths := []string{"/bin", "/sbin", "/usr/bin", "/usr/sbin", "/usr/local/bin", "/usr/local/sbin"}

	if contains(defaultPaths, element) {
		printlnColor(colorMap, "GRAY", element)
	} else if strings.Contains(element, "/go/") || strings.HasSuffix(element, "/go") {
		printlnColor(colorMap, "CYAN", element)
	} else if strings.Contains(element, "venv") || strings.Contains(element, "conda") || strings.Contains(element, "python") || strings.Contains(element, "Python") || strings.Contains(element, "py") {
		printlnColor(colorMap, "GREEN", element)
	} else if strings.Contains(element, "homebrew") || strings.Contains(element, "linuxbrew") {
		printlnColor(colorMap, "BROWN", element)
	} else if strings.Contains(element, "/.cargo/") {
		printlnColor(colorMap, "ORANGE", element)
	} else {
		printlnColor(colorMap, "BOLD", element)
	}
}

func main() {
	allPtr := flag.Bool("all", false, "Include duplicate paths from $PATH variable. (default false)")
	colorizePtr := flag.String("colorize", "8bit", "Choose how to color the output [8bit, 3bit, none].")
	flag.Parse()
	allArg := *allPtr
	colorizeArg := *colorizePtr

	stringPaths := os.Getenv("PATH")
	paths := strings.Split(stringPaths, ":")

	if !allArg {
		paths = unique(paths)
	}

	var colorMap map[string]string
	if colorizeArg != "8bit" && colorizeArg != "3bit" && colorizeArg != "none" {
		fmt.Println("--colorize expects one of [8bit, 3bit, none].")
		os.Exit(1)
	} else if colorizeArg == "8bit" {
		colorMap = colors8bit
	} else if colorizeArg == "3bit" {
		colorMap = colors3bit
	}

	for _, element := range paths {
		colorizePath(colorMap, element)
	}
}
