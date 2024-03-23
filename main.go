package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	fmt.Println(os.Environ)
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "лох?" {

		//Call the KuteGo API and retrieve our cute Dr Who Gopher
		//response, err := http.Get(KuteGoAPIURL + "/gopher/" + "dr-who")
		//if err != nil {
		//	fmt.Println(err)
		//}
		//defer response.Body.Close()
		//
		//if response.StatusCode == 200 {
		//	_, err = s.ChannelFileSend(m.ChannelID, "dr-who.png", response.Body)
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//} else {
		//	fmt.Println("Error: Can't get dr-who Gopher! :-(")
		_, err := s.ChannelMessageSend(m.ChannelID, "Сам лох")
		_, err2 := s.ChannelMessageSend(m.ChannelID, "<:6880pepe17:1063153862409736303>")
		if err != nil {
			return
		}
		if err2 != nil {
			return
		}
	}
}
