package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/spetr/tdu/units"
)

func init() {
	var (
		err error
	)

	flag.Parse()

	// Parse and check start path
	confStartPath = path.Clean(*flagStartPath)
	if _, err = os.Stat(confStartPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error checking start path %s: %s\n", confStartPath, err)
		os.Exit(1)
	}

	// Parse and check change of directory size
	confChangeDirSize, err = units.DataSizeParse(*flagChangeDirSize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing csize size: %s\n", err)
		os.Exit(1)
	}
	if confChangeDirSize <= 0 {
		fmt.Fprint(os.Stderr, "Error parsing csize size: can not be 0 or less\n")
		os.Exit(1)
	}

	// Parse and check time
	confTimeAfter = time.Now().Add(-*flagFTime)
	confTimeBefore = time.Now().Add(-*flagTTime)

	if *flagStartPath == "" {
		fmt.Fprintf(os.Stderr, "TDU (Time based Disk Usage) version: %s\n\n", version)
		flag.Usage()
		os.Exit(1)
	}
}
