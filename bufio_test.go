package notes

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestIO(t *testing.T) {
	var p string
	p, _ = os.Getwd()
	file, err := os.Open(filepath.Join(p, "sample.txt"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("Error reading file:", err)
			return
		}

		fmt.Print(line)
	}

	/*
		// src is the input that we want to tokenize.
		src := []byte("cos(x) + 1i*sin(x) // Euler")

		// Initialize the scanner.
		var s scanner.Scanner
		fset := token.NewFileSet()                      // positions are relative to fset
		file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
		s.Init(file, src, nil /* no error handler */ /*, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)

	}*/
}
