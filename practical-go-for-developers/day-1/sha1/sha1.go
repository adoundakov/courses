package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close()
	var r io.Reader = file

	// If we want to use this to decode both gzip and regular files,
	// we can define a variable to hold an io.Reader, but that can come from either
	// the file or the gzip reader.
	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("error unzipping file: %w", err)
		}
		defer gz.Close()
		r = gz
	}

	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("error generating shasum: %w", err)
	}

	sum := w.Sum(nil)

	return fmt.Sprintf("%x", sum), nil
}

func main() {
	shaSum, err := sha1Sum("./hello.log.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Printf("Gzip Sha Sum: %s | Is Correct: %t", shaSum, shaSum == "22596363b3de40b06f981fb85d82312e8c0ed511")
	fmt.Println()

	shaSum, err = sha1Sum("./hello.log")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Printf("Gzip Sha Sum: %s | Is Correct: %t", shaSum, shaSum == "22596363b3de40b06f981fb85d82312e8c0ed511")
}
