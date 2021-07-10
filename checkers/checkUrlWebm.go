package checkers

import "regexp"

//смотрим есть ли url в сообщении
func UrlWebm(s string) string {
	r := string(regexp.MustCompile(`\s*(https:.+)\.webm\s*`).Find([]byte(s)))
	return string(regexp.MustCompile(`(https:.+)\.webm`).Find([]byte(r)))
}
