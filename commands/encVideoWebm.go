package commands

import (
	"botFFMPEG/actions"
	"botFFMPEG/checkers"
	"fmt"
	"time"

	"regexp"

	"github.com/bwmarrin/discordgo"
)

// var freeSlot int
// var queue []string
var typeMedia = regexp.MustCompile(`\.webm`)

func EncHandlerWebm(s *discordgo.Session, m *discordgo.MessageCreate) {
	//отсекаем детект сообщений от самого бота
	if m.Author.ID == s.State.User.ID {
		return
	}
	// определяем url и названия вложенных файлов
	var urlAtch string
	var nameAtch string
	for _, atch := range m.Attachments {
		urlAtch = atch.URL
		nameAtch = atch.Filename
	}
	urlEmb := checkers.UrlWebm(m.Content)
	nameEmb := checkers.NameUrl(urlEmb)
	// определяем webm в файле или в ссылке
	if typeMedia.MatchString(nameAtch) || typeMedia.MatchString(nameEmb) {
		// даем знать в чат, что webm определен и записываем айдишку этого меседжа
		s.ChannelMessageSend(m.ChannelID, "WEBM detected. Wait few moments... :clapper:")
		massageWaitId := actions.LastIdMessageWebm
		//очередь для обработки файлов
		actions.GetQueue(m.ID)
		//замеряем старт работы с файлом
		t := time.Now()
		// определяем тип(ссылка или файл), от этого определяем вложенное сообщение
		var url string
		var name string
		var message string
		content := actions.NewMesContent(s, m)
		if typeMedia.MatchString(nameAtch) {
			url = urlAtch
			name = nameAtch
			message = "(webm)" + m.Author.Username + ": " + content + "||оригинал: " + url + "||"
		} else {
			url = urlEmb
			name = nameEmb
			message = "(webm)" + m.Author.Username + ": " + content
		}
		//логи
		fmt.Println(actions.LogMessage(name, "start", checkers.CheckMbUrl(url), t, s, m))
		// преобразуем название в mp4 для ffmpeg
		newName := name[:(len(name)-5)] + ".mp4"
		// используем ffmpeg в системе, для конвертации
		ffmpeg := "ffmpeg -fflags +genpts -i " + url + " -r 24 videos/" + m.ID + "/" + newName
		path := "videos/" + m.ID
		if !actions.UseFfmpeg(ffmpeg, path) {
			//в случае фейла делаем ответ с типом операции
			typeOperation := "webm to mp4"
			actions.MessageBadAnswer(massageWaitId, typeOperation, s, m)
			//логи
			fmt.Println(actions.LogMessage(name, "fail", "0", t, s, m))
		} else {
			//логи
			fmt.Println(actions.LogMessage(name, "complete", checkers.CheckMbFile(m.ID, newName), t, s, m))
			// отправляем ответ с конвертированным файлом
			actions.MessageAnswer(newName, massageWaitId, message, s, m)
		}
		//выходим из очереди файлов
		actions.FreeSlot = actions.FreeSlot - 1
	}

}
