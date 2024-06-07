package main

import (
	"fmt"
	"routinesChannels/download"
	"sync"
)

func main() {

	download := []download.Download{
		{"https://example.com/file1.zip", "file1.zip"},
		{"https://example.com/file2.zip", "file2.zip"},
		{"https://example.com/file3.zip", "file3.zip"},
	}

	var wg sync.WaitGroup

	ch := make(chan string, len(download))

	for _, download := range download {
		wg.Add(1)
		go download.DownloadFile(&wg, ch)
	}
	wg.Wait()

	for i := 0; i < len(download); i++ {
		downloadedFile := <-ch
		fmt.Println("Downloaded file:", downloadedFile)
	}

	fmt.Println("Download Complete.")
}
