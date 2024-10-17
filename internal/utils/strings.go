package utils

import (
	"bytes"
	"strconv"
	"strings"
	"unicode/utf8"
)

func TruncateString(s string, length int) string {
	if utf8.RuneCountInString(s) <= length {
		return s
	}

	// Convert the string to a slice of runes
	runes := []rune(s)
	return string(runes[:length])
}

func ToInt(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return id
}

func ToIntDefault(s string, i int) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		return i
	}
	return id
}

func ToUInt(s string) uint {
	id, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint(id)
}

// ToIntSlice 将逗号分隔的数字字符串转换为 []int 切片
func ToIntSlice(ids string) []int {
	// 分割字符串
	strIds := strings.Split(ids, ",")
	var intIds []int

	// 遍历字符串数组并转换为整数
	for _, strId := range strIds {
		strId = strings.TrimSpace(strId) // 移除前后空格
		if id, err := strconv.Atoi(strId); err == nil {
			intIds = append(intIds, id)
		}
	}

	return intIds
}
func ToInt64Slice(ids string) []int64 {
	// 分割字符串
	strIds := strings.Split(ids, ",")
	var intIds []int64

	// 遍历字符串数组并转换为整数
	for _, strId := range strIds {
		strId = strings.TrimSpace(strId) // 移除前后空格
		if id, err := strconv.ParseInt(strId, 10, 64); err == nil {
			intIds = append(intIds, id)
		}
	}

	return intIds
}
func ToInt64(str string) int64 {
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return id
}
func IsTextFile(ob []byte) (bool, error) {

	n := len(ob)
	if n > 1024 {
		n = 1024
	}
	// 检查是否包含非文本字符
	if !utf8.Valid(ob[:n]) {
		return false, nil
	}

	// 检查是否包含空字节（\x00），空字节通常代表二进制文件
	if bytes.Contains(ob[:n], []byte{0}) {
		return false, nil
	}

	return true, nil
}
