package commands

import (
	"botFFMPEG/actions"
	"botFFMPEG/checkers"
	"fmt"
	"time"

	"regexp"

	"github.com/bwmarrin/discordgo"
)

func EncHandlerMain(s *discordgo.Session, m *discordgo.MessageCreate) {
	//отсекаем детект сообщений от самого бота
	if m.Author.ID == s.State.User.ID {
		return
	}
	// определяем url и названия вложенных файлов
	var urlFile string
	var fileName string
	for _, atch := range m.Attachments {
		urlFile = atch.URL
		fileName = atch.Filename
	}
	// определяем hevc в mp4 файлах
	if regexp.MustCompile(`\.mp4`).MatchString(fileName) && checkers.CheckMain(urlFile) {
		// даем знать в чат, что main определен и записываем айдишку этого меседжа
		s.ChannelMessageSend(m.ChannelID, "h264 (main) detected. Wait few moments... :clapper:")
		massageWaitId := actions.LastIdMessageMain
		//очередь для обработки файлов
		actions.GetQueue(m.ID)
		//замеряем старт работы с файлом
		t := time.Now()
		//логи
		fmt.Println(actions.LogMessage(fileName, "start", checkers.CheckMbUrl(urlFile), t, s, m))
		// указываем отправителя и его сообщение(обработанное) с вложенным видео
		content := actions.NewMesContent(s, m)
		message := "(Main)" + m.Author.Username + ": " + content
		// используем ffmpeg в системе, для конвертации
		ffmpeg := "ffmpeg -i " + urlFile + " -c:v libx264 -profile:v high -c:a copy videos/" + m.ID + "/" + fileName
		path := "videos/" + m.ID
		if !actions.UseFfmpeg(ffmpeg, path) {
			//в случае фейла делаем ответ с типом операции
			typeOperation := "h.264 (main) to h.264 (high)"
			actions.MessageBadAnswer(massageWaitId, typeOperation, s, m)
			// логи
			fmt.Println(actions.LogMessage(fileName, "fail", "0", t, s, m))
		} else {
			//логи
			fmt.Println(actions.LogMessage(fileName, "complete", checkers.CheckMbFile(m.ID, fileName), t, s, m))
			// отправляем ответ с конвертированным файлом
			actions.MessageAnswer(fileName, massageWaitId, message, s, m)
		}
		//выходим из очереди файлов
		actions.FreeSlot = actions.FreeSlot - 1
	}

}
