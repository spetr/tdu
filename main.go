package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/spetr/tdu/units"
)

var (
	version = "dev"

	flagStartPath     = flag.String("path", "", "start path for walking")
	flagFTime         = flag.Duration("ftime", time.Hour*24, "from time in hours")
	flagTTime         = flag.Duration("ttime", time.Hour*0, "to time in hours")
	flagChangeDirSize = flag.String("csize", "100MB", "size of changes in one directory")
	flagCSV           = flag.String("csv", "", "output in csv format (optional)")

	confStartPath     = ""
	confChangeDirSize = int64(0)
	confTimeAfter     time.Time
	confTimeBefore    time.Time
	csvFile           *os.File
)

func main() {
	fmt.Fprintf(os.Stderr, "TDU (Time based Disk Usage) version: %s\n", version)
	fmt.Fprint(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Start path: %s\n", *flagStartPath)
	fmt.Fprintf(os.Stderr, "Min directory change size: %s\n", units.DataSizeHuman(confChangeDirSize, 2))
	fmt.Fprintf(os.Stderr, "Count only files created/modified after: %s\n", confTimeAfter.Format(time.RFC3339))
	fmt.Fprintf(os.Stderr, "Count only files created/modified before: %s\n", confTimeBefore.Format(time.RFC3339))
	fmt.Fprint(os.Stderr, "\n")

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
			if itemInfo.ModTime().After(confTimeAfter) && itemInfo.ModTime().Before(confTimeBefore) {
				sizeSum += itemInfo.Size()
			}
		}
	}

	// Report if sizeSum is greater than min size
	if sizeSum > confChangeDirSize {
		fmt.Printf("%s %s\n", currentPath, units.DataSizeHuman(sizeSum, 1))
		// Write to CSV file
		if csvFile != nil {
			fmt.Fprintf(csvFile, "%s,%d\n", currentPath, sizeSum)
		}
	}

	// Process subdirectories
	for i := range directories {
		processDirectory(directories[i])
	}
}
