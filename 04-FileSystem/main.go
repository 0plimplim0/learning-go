package main

import (
	"fmt"
	"strings"
	"sync"
)

func encrypt(data string) string {
	runes := []rune(data)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type Processer interface {
	Process() (string, error)
}

type Compressor struct {
	data string
}

func (c Compressor) Process() (string, error) {
	compressed := strings.ToUpper(c.data)
	return compressed, nil
}

type Encryptor struct {
	data string
}

func (e Encryptor) Process() (string, error) {
	encrypted := encrypt(e.data)
	return encrypted, nil
}

func AsyncRun(process Processer, wg *sync.WaitGroup) {
	defer wg.Done()
	result, err := process.Process()
	if err != nil {
		fmt.Println("There was an error processing the data...")
		return
	}
	fmt.Printf("Process complete: %s\n", result)
}

func main() {
	var wg sync.WaitGroup
	tasks := []Processer{
		Compressor{data: "Hello world"},
		Encryptor{data: "ultrasecret"},
	}
	for _, v := range tasks {
		wg.Add(1)
		go AsyncRun(v, &wg)
	}
	wg.Wait()
}
