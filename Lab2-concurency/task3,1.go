package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(filename string, sizeMB int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Downloading %s (%dMB)...\n", filename, sizeMB)
	time.Sleep(time.Duration(sizeMB*100) * time.Millisecond)
	fmt.Printf("✓ %s complete!\n", filename)
}

func main() {
	fmt.Println("=== File Downloader ===")
	start := time.Now()

	files := map[string]int{
		"video.mp4":   8,
		"song.mp3":    4,
		"photo.jpg":   2,
		"doc.pdf":     5,
		"archive.zip": 6,
	}

	var wg sync.WaitGroup
	for name, size := range files {
		wg.Add(1)
		go downloadFile(name, size, &wg)
	}

	wg.Wait()
	fmt.Printf("\n✓ All downloads complete! (%s)\n", time.Since(start))
}
