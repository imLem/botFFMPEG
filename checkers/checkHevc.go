package checkers

import (
	"log"
	"os/exec"
	"strings"
)

//чекаем hevc кодек в mp4 файле
func CheckHevc(s string) bool {
	encType := "ffprobe -v error -select_streams v:0 -show_entries stream=codec_name -of default=noprint_wrappers=1:nokey=1 " + s

	args := strings.Split(encType, " ")

	cmd := exec.Command(args[0], args[1:]...)

	result, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Running FFmpeg failed: %v", err)
	}
	if string(result) == string([]byte{104, 101, 118, 99, 13, 10}) {
		return true
	} else {
		return false
	}
}
