package commands

import (
	"botFFMPEG/actions"
	"botFFMPEG/checkers"
	"fmt"

	"regexp"

	"github.com/bwmarrin/discordgo"
)

func EncHandlerHevc(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	if regexp.MustCompile(`.mp4`).MatchString(fileName) && checkers.CheckHevc(urlFile) {
		// даем знать в чат, что hevc определен и записываем айдишку этого меседжа
		s.ChannelMessageSend(m.ChannelID, "HEVC detected. Wait few moments... :clapper:")
		massageWaitId := actions.LastIdMessageHevc
		//очередь для обработки файлов
		actions.GetQueue(m.ID)
		// указываем отправителя и его сообщение(обработанное) с вложенным видео
		content := actions.NewMesContent(s, m)
		message := m.Author.Username + ": " + content
		// используем ffmpeg в системе, для конвертации
		ffmpeg := "ffmpeg -i " + urlFile + " -vcodec libx264 -acodec aac videos/" + m.ID + "/" + fileName
		path := "videos/" + m.ID
		if !actions.UseFfmpeg(ffmpeg, path) {
			//в случае фейла делаем ответ с типом операции
			typeOperation := "hevc to h.264"
			actions.MessageBadAnswer(massageWaitId, typeOperation, s, m)
		} else {
			// отправляем ответ с конвертированным файлом
			actions.MessageAnswer(fileName, massageWaitId, message, s, m)
			fmt.Println(fileName + " кейс HEVC завершен")
		}
		//выходим из очереди файлов
		actions.FreeSlot = actions.FreeSlot - 1
	}

}
