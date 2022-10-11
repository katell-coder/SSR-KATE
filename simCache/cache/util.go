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
	//默认大小 100MB
	re, _ := regexp.Compile("[0-9]+")
	unist := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unist, "", 1), 10, 64)
	unit := strings.ToUpper(unist)
	var byteNum int64 = 0

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
		num = 0
	}
	if num == 0 {
		log.Println("parseSize 仅仅支持 B,KB,MB,GB,PB")
		num = 100
		byteNum = 100 * MB
		unist = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unist

	return byteNum, sizeStr
}

func GetValSize(val interface{}) int64 {
	marshal, err := json.Marshal(val)
	if err != nil {
		log.Println("GetValSize err", marshal)
	}

	return int64(len(marshal))
}
