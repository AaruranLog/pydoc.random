package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"io/ioutil"
	"sync"
	"strings"
	"flag"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readMapToCsv(filename string) map[string]string{
	m := make(map[string]string)
	in, err := os.Open(filename)
	check(err)
	r := csv.NewReader(in)
	for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			m[record[0]] = record[1]
	}
	fmt.Println(m)
	return m
}

func encodeText(sourceFilename, targetFilename string, m *map[string]string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	if strings.HasPrefix(sourceFilename, ".") {
		fmt.Println("Hidden file, skipping")
		return
	}
	in, err := os.Open(sourceFilename)
	check(err)

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanRunes)

	t, err := os.Create(targetFilename)
	check(err)
	defer t.Close()
	// result := make(map[string] int)
	fmt.Println("encoding:", sourceFilename)
	for s.Scan() {
		c := s.Text()
		newC, exists := (*m)[c]
		if ! exists {
			newC = (*m)[" "]
		}
		t.WriteString(newC + " ")
	}
	fmt.Println("finished encoding:", sourceFilename)
	return
}


func main() {
	rootDirectoryPtr := flag.String("dir", "data/raw/example_text_files/", "The directory holding all .txt files to count")
	targetDirectoryPtr := flag.String("target", "data/interim/cleaned_examples/", "The directoy which will hold the encoded .txt files")
	flag.Parse()
	rootDirectory := *rootDirectoryPtr
	targetDirectory := *targetDirectoryPtr
	c, err := ioutil.ReadDir(rootDirectory)
	check(err)

	var wg sync.WaitGroup
	numberOfFiles := len(c)
	wg.Add(numberOfFiles)
	// Use a buffered channel, as it is non-blocking
	converter := readMapToCsv("src/features/char_int_map.csv")

	for _, v := range c {
		sourceFilename := rootDirectory + v.Name()
		targetFilename := targetDirectory + v.Name()
		fmt.Println("sourceFilename:", sourceFilename)
		fmt.Println("targetFilename:", targetFilename)
		go encodeText(sourceFilename,targetFilename, &converter, &wg)
	}
	wg.Wait()
}
