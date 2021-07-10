package actions

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func NewMesContent(s *discordgo.Session, m *discordgo.MessageCreate) string {

	newCont := m.Content

	r := regexp.MustCompile(`<@&[0-9]+>`)
	r2 := regexp.MustCompile(`<@![0-9]+>`)
	r3 := string(regexp.MustCompile(`\s*(https:.+)\.webm\s*`).Find([]byte(newCont)))
	url := string(regexp.MustCompile(`(https:.+)\.webm`).Find([]byte(r3)))

	// r3 := regexp.MustCompile(`[0-9]+`)
	if url != "" {
		newCont = strings.ReplaceAll(newCont, url, "<"+url+">")
	}
	channelData, _ := s.Channel(m.ChannelID)
	guildID := channelData.GuildID
	guildData, _ := s.Guild(guildID)

	newCont = (r.ReplaceAllStringFunc(newCont, func(v string) string {

		for _, g := range guildData.Roles {
			if "<@&"+g.ID+">" == v {
				return g.Name
			}
		}
		// func (s *State) Role(guildID, roleID string) (*Role, error) {}
		return ""

	}))

	newCont = (r2.ReplaceAllStringFunc(newCont, func(v string) string {
		id := strings.ReplaceAll(v, "<@!", "")
		id = strings.ReplaceAll(id, ">", "")
		member, _ := s.GuildMember(guildID, id)
		user := member.User
		name := user.Username
		return name
	}))

	newCont = strings.ReplaceAll(newCont, "@", ":warning:")

	return newCont

}
