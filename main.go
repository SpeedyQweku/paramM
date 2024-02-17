package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	PatFile   string
	StdinData []string
	PatData   []string
	Err       error
}

func main() {
	var cfg Config

	flag.StringVar(&cfg.PatFile, "l", "", "file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if cfg.PatFile != "" {
		cfg.PatData, cfg.Err = readFile(cfg.PatFile)
		if cfg.Err != nil {
			log.Fatalf("Error reading pattern file: %v", cfg.Err)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

	if stdinIsPipe() {
		cfg.Err = readStdin(&cfg)
		if cfg.Err != nil {
			log.Fatalf("Error reading stdin: %v", cfg.Err)
		}
	} else {
		fmt.Println("stdin is not from a pipe")
	}

	for _, stdata := range cfg.StdinData {
		processLine(stdata, &cfg)
	}
}

func processLine(line string, cfg *Config) {
	for _, pattern := range cfg.PatData {
		trimmedPattern := strings.Trim(pattern, " ")
		if !strings.HasSuffix(trimmedPattern, "=") {
			trimmedPattern = trimmedPattern + "="
		}
		if strings.Contains(line, "?"+trimmedPattern) || strings.Contains(line, "&"+trimmedPattern) {
			fmt.Printf("%s : %s\n", trimmedPattern, line)
			break
		}
	}
}

func readStdin(cfg *Config) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		cfg.StdinData = append(cfg.StdinData, line)
	}
	return scanner.Err()
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func stdinIsPipe() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
