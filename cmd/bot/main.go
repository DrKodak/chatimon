package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const BOT_PREFIX string = "!chatimon"

func main() {
	godotenv.Load("../../.env")
	token := os.Getenv("BOT_TOKEN")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("ERROR Creating Discord Session: ", err)
	}

	// Register the messageCreate func as a callback for the MessageCreate event
	session.AddHandler(messageCreate);	

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	
	err = session.Open()
	if err != nil {
		log.Fatal("ERROR Could not open session")
	}
	
	defer session.Close()

	fmt.Println("Bot is online")

	// Create a channel to sit open and listen for a system interrupt.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// This function will be called every time a new message is created in a channel that the bot has access to
func messageCreate(session *discordgo.Session, m *discordgo.MessageCreate) {
	if session.State.User.ID == m.Author.ID {
		return
	}

	args := strings.Split(m.Content, " ")

	if args[0] != BOT_PREFIX {
		return
	}
	
	if args[1] == "showMon" {
		// TODO: Add logic t talk to server and fetch Mon state
		session.ChannelMessageSend(m.ChannelID, "Your Mon is happy!")
	}

	if args[1] == "pet" {
		file, err := os.Open("../../images/agumon.gif")
		if err != nil {
			log.Printf("Could not open GIF: %v", err)
		}
		defer file.Close()

		_, err = session.ChannelFileSend(m.ChannelID, "digimon.gif", file)
		if err != nil {
			log.Printf("Could not send GIF: %v", err)
		}
	}
}

// slashCommandHandler handles slash command interactions
func slashCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
}