package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	PatFile string
	UrlFile string
	Pattern string
}

func main() {
	var cfg Config
	fmt.Println("Hello paramM")

	flag.StringVar(&cfg.PatFile, "p", "", "parameter file")
	flag.StringVar(&cfg.UrlFile, "l", "", "url file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if cfg.PatFile == "" && cfg.UrlFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	patData, err := readFile(cfg.PatFile)
	if err != nil {
		log.Fatalf("Error reading pattern file: %v", err)
	}
	var cmd *exec.Cmd

	for _, pattern := range patData {
		trimmedPattern := strings.Trim(pattern, " ")
		cfg.Pattern = fmt.Sprintf("\\?%s=|&%s=", trimmedPattern, trimmedPattern)
		cmd = exec.Command("grep", "-iE", cfg.Pattern, cfg.UrlFile)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
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