package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Kori-Sama/compiler-go/lexer"
)

const (
	OUTPUT_SUFFIX = ".out"
)

func main() {
	inputPath, outputPath := parse_args()

	fmt.Printf("Input path: %s\n", inputPath)
	fmt.Printf("Output path: %s\n", outputPath)

	input := read_file(inputPath)

	lexer := lexer.NewLexer(input)

	tokens := lexer.ParseAll()

	if lexer.Err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", lexer.Err)
		os.Exit(1)
	}

	output := ""
	for _, token := range tokens {
		output += fmt.Sprintf("%v: %s\n", token.Kind, token.Literal)
	}

	fmt.Print(output)
}

func read_file(path string) *string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}

	str := string(data)
	return &str
}

func parse_args() (inputPath string, outputPath string) {
	program := os.Args[0]

	if len(os.Args) == 0 {
		usage(os.Stderr, program)
		os.Exit(1)
	}
	idx := 1
	for idx < len(os.Args) {
		switch os.Args[idx] {
		case "-o":
			idx++
			if idx >= len(os.Args) {
				fmt.Fprintf(os.Stderr, "ERROR: Missing output path\n")
				os.Exit(1)
			}
			outputPath = os.Args[idx]
		case "-h":
			usage(os.Stdout, program)
			os.Exit(0)
		default:
			inputPath = os.Args[idx]
		}
		idx++
	}

	if inputPath == "" {
		fmt.Fprintf(os.Stderr, "ERROR: Missing input path\n")
		os.Exit(1)
	}

	if outputPath == "" {
		outputPath = inputPath + OUTPUT_SUFFIX
	}

	return inputPath, outputPath
}

func usage(w io.Writer, program string) {
	fmt.Fprintf(w, "Usage: %s [options] <input>\n", program)
	fmt.Fprintf(w, "Options:\n")
	fmt.Fprintf(w, "    -o <output>     Provide output path\n")
	fmt.Fprintf(w, "    -h              Show this help message\n")
}
