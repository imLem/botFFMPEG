package checkers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//чекаем hevc кодек в mp4 файле
func CheckMbUrl(url string) string {
	var clt = &http.Client{}

	rsp, err := clt.Head(url)
	if err != nil {
		fmt.Println("HEAD request failed", err)
		return "0"
	} else {
		// по-хорошему, тут надо обработать статус запроса
		return fmt.Sprintf("%4.2f", float64(rsp.ContentLength)/1048576)
	}

}
func CheckMbFile(messageId, fileName string) string {
	file, err := os.Open("videos/" + messageId + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	//получем размер, конвертированного файла
	fi, _ := file.Stat()
	sizeFile := float64(fi.Size()) / 1048576
	return fmt.Sprintf("%4.2f", sizeFile)
}
