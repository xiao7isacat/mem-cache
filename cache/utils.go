package cache

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	var byteNum int64 = 0
	//默认大小为100mb
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		byteNum = 0
	}

	if num == 0 {
		log.Println("Parse size 仅支持 B KB MB GB TB PB，默认设置为100MB")
		num = 100
		byteNum = 100 * MB
		unit = "MB"
	}
	sizestr := strconv.FormatInt(num, 10) + unit
	return byteNum, sizestr
}

func GetValSize(val interface{}) int64 {
	valBytes, _ := json.Marshal(val)
	size := int64(len(valBytes))

	return size
}
