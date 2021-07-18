package checkers

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

//чекаем h264 (main) кодек в mp4 файле
func CheckMain(s string) bool {
	encType := "ffprobe " + s

	args := strings.Split(encType, " ")

	cmd := exec.Command(args[0], args[1:]...)

	result, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Running FFmpeg failed: %v", err)
	}

	if regexp.MustCompile(`h264 \S(Main)\S`).MatchString(string(result)) {
		return true

	} else {
		return false

	}
}
