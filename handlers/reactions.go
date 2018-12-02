package handlers

import (
	"fmt"
	"strings"

	"github.com/cactauz/gobot" 
	"github.com/bwmarrin/discordgo"
) 

const ()

func init() {
	gobot.Global.AddMessageHandler(NewContainsHandler(gobot.BOT_ID, fingerGunsReplyHandler))
	gobot.Global.AddMessageHandler(NewContainsHandler("lmao", lmaoHandler))
	gobot.Global.AddMessageHandler(NewContainsHandler("nice", niceHandler))
	gobot.Global.AddMessageHandler(NewContainsHandler(" b ", bHandler))
	gobot.Global.AddMessageHandler(NewPrefixHandler("F", fHandler))
}

func fHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "F" {
		return
	}

	c := rune(127462 + int('f') - 97)
		
	err := s.MessageReactionAdd(m.ChannelID, m.ID, fmt.Sprintf("%c", c))
	if err != nil {
		fmt.Println(err)
	}
} 

func lmaoHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	id := m.ID

	prev, _ := s.ChannelMessages(m.ChannelID, 1, m.ID, "", "")

	// if the message is this exactly react to previous message
	if len(prev) > 0 && m.Content == "lmao" {
		id = prev[0].ID
	}

	if !gobot.GlobalTimeout(m.ChannelID) {
		return
	}

	res := []rune{'l', 'm', 'a', 'o', '😂', '😹'}
	for _, c := range res {
		if c <= 'z' {
			c = rune(127462 + int(c) - 97)
		}
		err := s.MessageReactionAdd(m.ChannelID, id, fmt.Sprintf("%c", c))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func niceHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !gobot.GlobalTimeout(m.ChannelID) {
		return
	}

	err := s.MessageReactionAdd(m.ChannelID, m.ID, gobot.NICE)
	if err != nil {
		fmt.Println(err)
	}
}

func bHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := s.MessageReactionAdd(m.ChannelID, m.ID, "🅱")
	if err != nil {
		fmt.Println(err)
	}
}

func frogsirenHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := s.MessageReactionAdd(m.ChannelID, m.ID, gobot.FROGSIREN)
	if err != nil {
		fmt.Println(err)
	}
}

func fingerGunsReplyHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	split := strings.SplitN(strings.TrimSpace(m.Content), " ", 2)
	message := ""
	if len(split) > 1 {
		message = split[1]
	}
	
	if strings.HasPrefix(message, "<") && strings.HasSuffix(message, ">") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%v> %v", m.Author.ID, message))
	} else {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%v> <%s>", m.Author.ID, gobot.FINGER_GUNS))
	}

	err := s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		println(err)
	}
}
