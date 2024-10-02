package units

import (
	"fmt"
	"strconv"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	PB = 1024 * TB
	EB = 1024 * PB
)

func DataSizeHuman(size int64, precision int) string {
	var (
		sizeStr string
		marker  = ""
	)

	// handle negative size
	if size < 0 {
		marker = "-"
		size = -size
	}

	// convert size to human-readable format
	switch {
	case size < KB:
		sizeStr = fmt.Sprintf("%s%dB", marker, size)
	case size < MB:
		sizeStr = fmt.Sprintf("%s%.*fKB", marker, precision, float64(size)/KB)
	case size < GB:
		sizeStr = fmt.Sprintf("%s%.*fMB", marker, precision, float64(size)/MB)
	case size < TB:
		sizeStr = fmt.Sprintf("%s%.*fGB", marker, precision, float64(size)/GB)
	case size < PB:
		sizeStr = fmt.Sprintf("%s%.*fTB", marker, precision, float64(size)/TB)
	case size < EB:
		sizeStr = fmt.Sprintf("%s%.*fPB", marker, precision, float64(size)/PB)
	default:
		sizeStr = fmt.Sprintf("%s%.*fEB", marker, precision, float64(size)/EB)
	}

	return sizeStr
}

func DataSizeParse(sizeStr string) (int64, error) {
	var (
		size      int64
		unit      string
		sizeBytes int64
		subZero   = false
	)

	// handle negative size
	if sizeStr[0] == '-' {
		subZero = true
		sizeStr = sizeStr[1:]
	}

	// check if sizeStr is numeric
	if i, err := strconv.Atoi(sizeStr); err == nil {
		size = int64(i)
		return size, nil
	}

	// check if sizeStr is in the format of "size unit"
	if _, err := fmt.Sscanf(sizeStr, "%d%s", &size, &unit); err != nil {
		return 0, err
	}

	// convert size to bytes
	switch unit {
	case "B":
		sizeBytes = size
	case "KB", "K":
		sizeBytes = size * KB
	case "MB", "M":
		sizeBytes = size * MB
	case "GB", "G":
		sizeBytes = size * GB
	case "TB", "T":
		sizeBytes = size * TB
	case "PB", "P":
		sizeBytes = size * PB
	case "EB", "E":
		sizeBytes = size * EB
	default:
		return 0, fmt.Errorf("invalid size unit: %s", unit)
	}

	if subZero {
		sizeBytes = -sizeBytes
	}

	return sizeBytes, nil
}

/*
// Parses the human-readable size string into the amount it represents.
func parseSize(sizeStr string, uMap unitMap) (int64, error) {
	// TODO: rewrite to use strings.Cut if there's a space
	// once Go < 1.18 is deprecated.
	sep := strings.LastIndexAny(sizeStr, "01234567890. ")
	if sep == -1 {
		// There should be at least a digit.
		return -1, fmt.Errorf("invalid size: '%s'", sizeStr)
	}
	var num, sfx string
	if sizeStr[sep] != ' ' {
		num = sizeStr[:sep+1]
		sfx = sizeStr[sep+1:]
	} else {
		// Omit the space separator.
		num = sizeStr[:sep]
		sfx = sizeStr[sep+1:]
	}

	size, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return -1, err
	}
	// Backward compatibility: reject negative sizes.
	if size < 0 {
		return -1, fmt.Errorf("invalid size: '%s'", sizeStr)
	}

	if len(sfx) == 0 {
		return int64(size), nil
	}

	// Process the suffix.

	if len(sfx) > 3 { // Too long.
		goto badSuffix
	}
	sfx = strings.ToLower(sfx)
	// Trivial case: b suffix.
	if sfx[0] == 'b' {
		if len(sfx) > 1 { // no extra characters allowed after b.
			goto badSuffix
		}
		return int64(size), nil
	}
	// A suffix from the map.
	if mul, ok := uMap[sfx[0]]; ok {
		size *= float64(mul)
	} else {
		goto badSuffix
	}

	// The suffix may have extra "b" or "ib" (e.g. KiB or MB).
	switch {
	case len(sfx) == 2 && sfx[1] != 'b':
		goto badSuffix
	case len(sfx) == 3 && sfx[1:] != "ib":
		goto badSuffix
	}

	return int64(size), nil

badSuffix:
	return -1, fmt.Errorf("invalid suffix: '%s'", sfx)
}
*/
