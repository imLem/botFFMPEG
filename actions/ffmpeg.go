package actions

import (
	"botFFMPEG/checkers"
	"log"
	"os"
	"os/exec"
	"strings"
)

// принимает команду для ffmpeg для конвертации файла
func UseFfmpeg(s, path string) bool {
	//проверка существует ли путь, если нет, то создаем его
	if !checkers.CheckFile(path) {
		os.Mkdir(path, 0777)
	}
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Running FFmpeg failed: %v", err)
		if checkers.CheckFile(path) {
			os.RemoveAll(path + "/")
		}
		return false
	}
	return true
}
