package checkers

import (
	"os"
)

//проверка есть ли файл или папка в каталоге
func CheckFile(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
