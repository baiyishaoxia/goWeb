package models

import (
	"fmt"
)

var (
	B       int64  = 1
	KB      int64  = 1024
	MB      int64  = 1024 * 1024
	GB      int64  = 1024 * 1024 * 1024
	TB      int64  = 1024 * 1024 * 1024 * 1024
	formatF string = "%8.2f"
)

func SizeFormat(sizeInBytes int64) string {
	switch {
	case B <= sizeInBytes && sizeInBytes < KB:
		return fmt.Sprintf("%dB", sizeInBytes)
	case KB <= sizeInBytes && sizeInBytes < MB:
		return fmt.Sprintf(formatF+"KB", float64(sizeInBytes)/float64(KB))
	case MB <= sizeInBytes && sizeInBytes < GB:
		return fmt.Sprintf(formatF+"MB", float64(sizeInBytes)/float64(MB))
	case GB <= sizeInBytes && sizeInBytes < TB:
		return fmt.Sprintf(formatF+"GB", float64(sizeInBytes)/float64(GB))
	case TB <= sizeInBytes:
		return fmt.Sprintf(formatF+"TB", float64(sizeInBytes)/float64(TB))
	default:
		return "0"
	}
}
