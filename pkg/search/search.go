package search

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"strings"
)

//Result описывает один результат поиска
type Result struct {
	//Фраза которую искали
	Phrase string
	// Целиком вся строка в которой нашли вхождение (без \n или  \r\n в конце)
	Line string
	// Номер строки ( начиная с 1 ) на которой нашли вхождение
	LineNum int64
	// Номер позиции ( начиная с 1) на которой нашли вхождение
	ColNum int64
}

// All ищет все вхождения phrase в текстовых файлах files
func All(ctx context.Context, phrase string, files []string) <-chan []Result {

	ch := make(chan []Result)

	var wg sync.WaitGroup

	for i := 0; i < len(files); i++ {
		file := files[i]

		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			file, _ := os.Open(fileName)
			defer file.Close()
			buf := bufio.NewReader(file)
				lines := make([]string,0)
			for  {
				line, err := buf.ReadString('\n')
				if err == io.EOF {
					break
				}
				lines = append(lines,line)
			}
			//lines := strings.Split(fileName, "\n")
			result := make([]Result, 0)
			for i2, line := range lines {
				if strings.Contains(line, phrase) {
					lineNum := int64(i2 + 1)
					colNum := int64(strings.Index(line, phrase))
					var found = Result{
						Phrase:  phrase,
						Line:    line,
						LineNum: lineNum,
						ColNum:  colNum,
					}
					result = append(result, found)
				}
			}

			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			default:
			}
			ch <- result
		}(file)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}


func Any(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)
	if len(files) ==0{
		close(ch)
		return ch
	}


	var wg sync.WaitGroup

	for i := 0; i < len(files); i++ {
		file := files[i]

		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			file, _ := os.Open(fileName)
			defer file.Close()
			buf := bufio.NewReader(file)
			lines := make([]string,0)
			for  {
				line, err := buf.ReadString('\n')
				if err == io.EOF {
					break
				}
				lines = append(lines,line)
			}
			//lines := strings.Split(fileName, "\n")
			result := make([]Result, 0)
			for i2, line := range lines {
				if strings.Contains(line, phrase) {
					lineNum := int64(i2 + 1)
					colNum := int64(strings.Index(line, phrase))
					var found = Result{
						Phrase:  phrase,
						Line:    line,
						LineNum: lineNum,
						ColNum:  colNum,
					}
					result = append(result, found)
				}
			}

			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			default:
			}
			ch <- result
		}(file)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
