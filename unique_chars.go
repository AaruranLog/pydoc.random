package main

import (
	"fmt"
	"os"
	// "strings"
	"bufio"
	"io/ioutil"
	"sync"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func countCharactersInFile(filename string, ch chan map[string]int, wg *sync.WaitGroup) {
	defer (*wg).Done()
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
			result[c] = 0
		} else {
			result[c] = count + 1
		}
	}
	ch <- result
	fmt.Println("finished with:", filename)
	return
}

func combine(s1, s2 *map[string]int) *map[string]int {
	for k := range *s2 {
		// c := s2[k]
		_, is_in_s1 := (*s1)[k]
		if is_in_s1 {
			(*s1)[k] += (*s2)[k]
		} else {
			(*s1)[k] = (*s2)[k]
		}
	}
	return s1
}

func main() {
	rootDirectory := "examples/" // TODO: change this to os.argv

	c, err := ioutil.ReadDir(rootDirectory)
	check(err)

	var wg sync.WaitGroup
	numberOfFiles := len(c)
	maps := make(chan map[string] int, numberOfFiles)


	wg.Add(numberOfFiles)
	for _, v := range c {
		filename := rootDirectory + v.Name()
		fmt.Println("filename:", filename)
		go countCharactersInFile(filename, maps, &wg)
	}
	counter := make(map[string] int)
	// counter := <-maps
	// fmt.Println(<-maps)
	// close(maps)
	// wg.Wait()

	for i := 1; i <= numberOfFiles; i++ {
		m := <- maps
		counter = *combine(&counter, &m)
	}
	// go func(){
	// 	for m := range maps {
	// 		counter = *combine(&counter, &m)
	// 	}
	// }()
	fmt.Println(counter)
}
