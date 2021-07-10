package actions

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var LastIdMessageWebm string
var LastIdMessageHevc string
var LastIdMessageFail string
var LastIdMessageSize string

// слушатель сообщения детекта от бота, чтобы запомнить его айдишку
func WaitMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// скипаем все сообщения не от нашего бота
	if m.Author.ID != s.State.User.ID {
		return
	}
	// если сообщение от нашего бота, то проверяем, что он написал и заносим айдишку нужного сообщения
	webmMessage := regexp.MustCompile(`WEBM`)
	hevcMessage := regexp.MustCompile(`HEVC`)
	failMessage := regexp.MustCompile(`fail`)
	size := regexp.MustCompile(`файл получился`)
	if webmMessage.MatchString(m.Content) {
		LastIdMessageWebm = m.ID
	}
	if hevcMessage.MatchString(m.Content) {
		LastIdMessageHevc = m.ID
	}
	if failMessage.MatchString(m.Content) {
		LastIdMessageFail = m.ID
	}
	if size.MatchString(m.Content) {
		LastIdMessageSize = m.ID
	}
}
