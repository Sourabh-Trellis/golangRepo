package download

import (
	"fmt"
	"sync"
	"time"
)

type Download struct {
	Url      string
	FileName string
}

func (d Download) DownloadFile(wg *sync.WaitGroup, ch chan<- string) {
	time.Sleep(time.Second * 2)
	fmt.Printf("Download complete: %s -> %s\n", d.Url, d.FileName)
	ch <- d.FileName
	wg.Done()
}


