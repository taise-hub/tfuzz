package tfuzz

import (
	"fmt"
	"sync"
	"net/http"
	"os"
	"bufio"
)

type Options struct {
	TargetUrl 	  string
	InputFile 	  string
}

type FuzzResult struct {
	StatusCode	int
	FuzzString	string
}

func (o *Options) ReadFile() []string {
	f, err := os.Open(o.InputFile)
	CheckErr(err, "Error ocuuerd while reading input file.")
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}


func (o *Options) StartFuzz(fuzz []string) {
	reqwg := &sync.WaitGroup{} //request wait group
	repwg := &sync.WaitGroup{} //reply wait group
	out := make(chan FuzzResult, len(fuzz))
	for _, s := range fuzz {
		reqwg.Add(1)
		go Request(o.TargetUrl, s,out, reqwg)
	}
	go func() {
		reqwg.Wait()
		close(out)
	}()
	
	repwg.Add(1)
	go func() {
		for r := range out {
			fmt.Println(r)
		}
		repwg.Done()
	}()
	repwg.Wait()
}

func Request(url, fuzzString string, out chan FuzzResult, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url +"/"+ fuzzString)
	CheckErr(err,"Error occuerd while sending HTTP request.", url, fuzzString)
	out <- FuzzResult{StatusCode: resp.StatusCode, FuzzString: fuzzString}
}