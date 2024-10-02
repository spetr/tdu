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
	flagCSV       = flag.String("csv", "", "output in csv format")
	csvFile       *os.File
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Fprintf(os.Stderr, "TDU version: %s\n", version)
	fmt.Fprintf(os.Stderr, "Modified before: %d\nMin increment size: %d\nStart path: %s\n", *flagTime, *flagMinSize, *flagStartPath)

	if *flagCSV != "" {
		fmt.Fprintf(os.Stderr, "Output in CSV format: %s\n", *flagCSV)
		var err error
		csvFile, err = os.Create(*flagCSV)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file %s: %s\n", *flagCSV, err)
			os.Exit(1)
		}
	}

	processDirectory(*flagStartPath)

	if csvFile != nil {
		csvFile.Close()
	}
}

func processDirectory(currentPath string) {
	// Read directory
	items, err := os.ReadDir(currentPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory %s: %s\n", currentPath, err)
		return
	}

	// Process items in directory
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

	// Report if sizeSum is greater than min size
	if sizeSum > int64(*flagMinSize)*1024*1024 {
		fmt.Printf("%s: %d MB\n", currentPath, sizeSum/1024/1024)
		// Write to CSV file
		if csvFile != nil {
			fmt.Fprintf(csvFile, "%s,%d\n", currentPath, sizeSum/1024/1024)
		}
	}

	// Process subdirectories
	for i := range directories {
		processDirectory(directories[i])
	}
}
