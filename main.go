package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"
)

var (
	version       = "dev"
	flagStartPath = flag.String("path", "./", "start path")
	flagTime      = flag.Int("time", 24, "time in hours")
	flagMinSize   = flag.Int("min", 100, "min size in MB")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Fprintf(os.Stderr, "TDU version: %s\n", version)
	fmt.Fprintf(os.Stderr, "Modified before: %d\nMin increment size: %d\nStart path: %s\n", *flagTime, *flagMinSize, *flagStartPath)
	processDirectory(*flagStartPath)
}

func processDirectory(currentPath string) {
	items, err := os.ReadDir(currentPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory %s: %s\n", currentPath, err)
		return
	}
	directories := []string{}
	sizeSum := int64(0)
	for i := range items {
		itemInfo, err := items[i].Info()
		if itemInfo.IsDir() {
			directories = append(directories, path.Join(currentPath, items[i].Name()))
		} else {
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file %s: %s\n", items[i].Name(), err)
				continue
			}
			if itemInfo.ModTime().Add(time.Hour * time.Duration(*flagTime)).After((time.Now())) {
				sizeSum += itemInfo.Size()
			}
		}
	}

	if sizeSum > int64(*flagMinSize)*1024*1024 {
		fmt.Printf("%s: %d MB\n", currentPath, sizeSum/1024/1024)
	}

	for _, directory := range directories {
		processDirectory(directory)
	}
}
