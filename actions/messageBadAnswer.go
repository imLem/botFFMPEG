package actions

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

//если не получается конвертировать файл делаем бед ответ
func MessageBadAnswer(massageWaitId, typeOperation string, s *discordgo.Session, m *discordgo.MessageCreate) {
	//удаляем сообщение с детектом
	s.ChannelMessageDelete(m.ChannelID, massageWaitId)
	//даем знать что произошел фейл
	s.ChannelMessageSend(m.ChannelID, typeOperation+" failed")
	massageWaitId2 := LastIdMessageFail
	//ожидаем 5 секунд и удаляем сообщение о фейле
	time.Sleep(5 * time.Second)
	s.ChannelMessageDelete(m.ChannelID, massageWaitId2)
}
