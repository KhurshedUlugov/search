package search

import (
	"bufio"
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

// Result описывает один результать поиска
type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}

// All ...
func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	wg := sync.WaitGroup{}
	ch := make(chan []Result)

	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(_ctx context.Context, index int, ch chan<- []Result) {
			defer wg.Done()
			result := []Result{}

			readFile, err := ioutil.ReadFile(files[index])
			if err != nil {
				log.Printf("Не удалось читать файл, %v", files[index])
			}

			lines := strings.Split(string(readFile), "\n")

			for in, line := range lines {
				cont := strings.Contains(line, phrase)
				if cont {
					colNum := strings.Index(line, phrase)

					res := Result{
						Phrase:  phrase,
						Line:    line,
						LineNum: int64(in + 1),
						ColNum:  int64(colNum + 1),
					}
					result = append(result, res)
				}
			}
			if len(result) > 0 {
				ch <- result
			}

		}(ctx, i, ch)

	}

	go func() {
		defer close(ch)
		wg.Wait()

	}()

	return ch
}

// Test ....
func Test(phrase string, files []string) []Result {
	result := []Result{}

	for i := 0; i < len(files); i++ {
		src, err := os.Open(files[i])
		if err != nil {
			log.Print(err)
		}

		defer func() {
			if cerr := src.Close(); cerr != nil {
				log.Print(cerr)
			}
		}()

		reader := bufio.NewReader(src)
		lineNum := int64(1)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				log.Print(line)
				break
			}

			if err != nil {
				log.Print(err)
				break
			}
			cantain := strings.Contains(line, "Shohin")
			//log.Print(cantain)
			if cantain {
				abs := strings.Index(line, "Shohin")
				//val := strings.Split(line, "\n")
				val := string(line[:len(line)-1])
				//val = val[:len(val)-1]
				log.Print("-----")
				log.Print(val)
				res := Result{
					Phrase:  phrase,
					Line:    val,
					LineNum: lineNum,
					ColNum:  int64(abs + 1),
				}
				result = append(result, res)
			}
			lineNum++
			//start := strings.HasPrefix(line, "Shohin")
			//log.Print(start)

		}
		//log.Print("----------------------------")
	}
	return result
}
