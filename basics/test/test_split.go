package test

import "strings"

// Split 将字符串 s 按照 sep 做切割，将结果保存到 result
func Split(s, sep string) (result []string) {
	idx := strings.Index(s, sep)
	for idx > -1 {
		result = append(result, s[:idx])
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	// 最后一段拼接到 result
	result = append(result, s)
	return
}
