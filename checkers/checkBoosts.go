package checkers

import (
	"github.com/bwmarrin/discordgo"
)

//проверка сколько бустов на сервере
func BoostCheck(s *discordgo.Session, m *discordgo.MessageCreate) int {
	channelData, _ := s.Channel(m.ChannelID)
	guildID := channelData.GuildID
	channelData2, _ := s.Guild(guildID)
	return channelData2.PremiumSubscriptionCount
}
