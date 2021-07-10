package checkers

import "regexp"

//узнаем названия файла в url
func NameUrl(s string) string {
	name := regexp.MustCompile(`[^\/]+.$`)
	return string(name.Find([]byte(s)))
}
