package search

import (
	"context"
	"fmt"
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
	result := make([]Result,0)

	wg := sync.WaitGroup{}
	wg.Add(1)

	for i := 0; i < len(files); i++ {
		file :=files[i]

		go func(file string) {
			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				wg.Done()
			}

			lines := strings.Split(file, "\n")

			for i2, line := range lines {

				if strings.Contains(line, phrase){
					lineNum := int64(i2 + 1)
					colNum := int64(strings.Index(line, phrase))
					var found  = Result{
						Phrase:  phrase,
						Line:    line,
						LineNum: lineNum,
						ColNum: colNum ,
					}
						result= append(result, found)
				}
			}
		}(file)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
