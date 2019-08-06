package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"io/ioutil"
	"sync"
	"strconv"
	"strings"
	"flag"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func countCharactersInFile(filename string, ch chan map[string]int, wg *sync.WaitGroup) {
	defer (*wg).Done()
	if strings.HasPrefix(filename, ".") {
		fmt.Println("Hidden file, skipping")
		return
	}

	in, err := os.Open(filename)
	check(err)

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanRunes)

	result := make(map[string] int)
	fmt.Println("counting:", filename)
	for s.Scan() {
		c := s.Text()
		count, exists := result[c]
		if ! exists {
			result[c] = 1
		} else {
			result[c] = count + 1
		}
	}
	ch <- result
	fmt.Println("done counting:", filename)
	return
}

func combine(s1, s2 *map[string]int) *map[string]int {
	for k := range *s2 {
		_, is_in_s1 := (*s1)[k]
		if is_in_s1 {
			(*s1)[k] += (*s2)[k]
		} else {
			(*s1)[k] = (*s2)[k]
		}
	}
	return s1
}

func writeMapToCsv(m *map[string]int, filename string) {
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	// Write head
	header := []string{"char", "count"}
	writer.Write(header)
	for k, v := range *m {
		line := []string{k, strconv.Itoa(v)}
		writer.Write(line)
	}

}

func main() {
	rootDirectoryPtr := flag.String("dir", "data/raw/example_text_files/", "The directory holding all .txt files to count")
	flag.Parse()
	rootDirectory := *rootDirectoryPtr

	c, err := ioutil.ReadDir(rootDirectory)
	check(err)

	var wg sync.WaitGroup
	numberOfFiles := len(c)
	wg.Add(numberOfFiles)
	// Use a buffered channel, as it is non-blocking
	maps := make(chan map[string] int, numberOfFiles)

	for _, v := range c {
		filename := rootDirectory + v.Name()
		fmt.Println("filename:", filename)
		go countCharactersInFile(filename, maps, &wg)
	}
	counter := make(map[string] int)

	for i := 1; i <= numberOfFiles; i++ {
		m := <- maps
		counter = *combine(&counter, &m)
	}

	fmt.Println(counter)
	writeMapToCsv(&counter, "counts.csv")
}
