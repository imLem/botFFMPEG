package actions

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func LogMessage(name, status, mb string, timeStart time.Time, s *discordgo.Session, m *discordgo.MessageCreate) string {
	channelData, _ := s.Channel(m.ChannelID)
	guildDate, _ := s.Guild(channelData.GuildID)
	t := timeStart
	duration := time.Since(t)
	ms := duration.Seconds()
	if status == "start" {
		end := " пошел в обработку"
		return "(" + guildDate.Name + " " + t.Format("15:04:05") + ")(" + m.Author.Username + "#" + m.Author.Discriminator + ") " + name + "(" + mb + " mb)" + end
	}
	if status == "fail" {
		end := " ошибка при конвертировании"
		return "(" + guildDate.Name + " " + t.Format("15:04:05") + ") " + name + end
	}
	if status == "complete" {
		end := " кейс завершен за "
		return "(" + guildDate.Name + " " + t.Format("15:04:05") + ") " + name + end + fmt.Sprintf("%.0f", ms) + " сек, новый размер файла: (" + mb + " mb)"
	}
	return ""
}
