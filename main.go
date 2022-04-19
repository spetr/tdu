package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"time"
)

var (
	flagStartPath = flag.String("path", "./", "start path")
	flagTime      = flag.Int("time", 24, "time in hours")
	flagMinSize   = flag.Int("min", 100, "min size in MB")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Printf("Start path: %s\n", *flagStartPath)
	processDirectory(*flagStartPath)
}

func processDirectory(currentPath string) {
	items, _ := ioutil.ReadDir(currentPath)
	directories := []string{}
	sizeSum := int64(0)
	for _, item := range items {
		if item.IsDir() {
			directories = append(directories, path.Join(currentPath, item.Name()))
		} else {
			if item.ModTime().Add(time.Hour * time.Duration(*flagTime)).After((time.Now())) {
				sizeSum += item.Size()
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
