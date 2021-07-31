package actions

import (
	"botFFMPEG/checkers"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

// функция формирует сообщение после детекта определенного формата
// fileName - имя файла, которое будет искаться в дирректории videos должно заканчиваться с форматом, например .mp4
// massageWaitId - айди сообщения которое было отправлено пользователю как реакция на детект формата, чтобы удалить его
// message сформированное тело сообщения, которое отправляется с файлом
func MessageAnswer(fileName, messageWaitId, message string, s *discordgo.Session, m *discordgo.MessageCreate) {
	// сначала определяем бусты на сервер, от этого зависит объем файла, который мы можем отправить
	boosts := checkers.BoostCheck(s, m)
	// открываем файл
	file, err := os.Open("videos/" + m.ID + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	//получем размер, конвертированного файла
	fi, _ := file.Stat()
	sizeFile := float64(fi.Size()) / 1048576
	// определяем какой размер файла мы можем отправить
	// если на сервере меньше 15 бустов, то отправить можно не более 8 мб
	// если 15 и более бустов и меньше 30, то отправить можно не более 50 мб
	// если 30 и более, то отправить можно не более 100 мб
	// если условия подошли, удаляем сообщение пользователя с файлом
	if (boosts < 15 && sizeFile < 8.0) || (boosts >= 15 && boosts < 30 && sizeFile < 50.0) || (boosts >= 30 && sizeFile < 100.0) {
		s.ChannelMessageDelete(m.ChannelID, messageWaitId)                 // если условия подошли, удаляем сообщение бота, что был детект
		s.ChannelMessageDelete(m.ChannelID, m.ID)                          // удаляем сообщение пользователя с файлом
		s.ChannelFileSendWithMessage(m.ChannelID, message, fileName, file) // отправляем конвертированный файл с сообщением
	}
	// если условия не подошли, то формируем сообщение неудачи, оставляя сообщение пользователя
	if (boosts < 15 && sizeFile > 8.0) || (boosts >= 15 && boosts < 30 && sizeFile > 50.0) || (boosts >= 30 && sizeFile > 100.0) {
		var sizeAtch string
		if sizeFile > 8.0 && sizeFile < 50.0 {
			sizeAtch = "8"
		}
		if sizeFile > 50.0 && sizeFile < 100.0 {
			sizeAtch = "50"
		}
		if sizeFile > 100.0 {
			sizeAtch = "100"
		}
		s.ChannelMessageDelete(m.ChannelID, messageWaitId)
		s.ChannelMessageSend(m.ChannelID, "файл получился больше "+sizeAtch+" мегабайт")
		messageWaitIdsize := LastIdMessageSize
		time.Sleep(5 * time.Second)
		s.ChannelMessageDelete(m.ChannelID, messageWaitIdsize)
		// file.Close()
		// os.RemoveAll("videos/" + m.ID + "/")
		// return
	}
	//закрываем работу с файлом и удаляем его
	file.Close()
	os.RemoveAll("videos/" + m.ID + "/")
}
