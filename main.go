package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// How to Running this go script
	// go main.py {discordToken} {channelID}

	if len(os.Args) < 3 {
		fmt.Println("Please not run this script directly without arguments")
		fmt.Println("Run Example: main.exe {discordToken} {channelID} {message}")
		fmt.Scanf("Press any button to leave..")
	}

	discordToken := os.Args[1]
	channelID := os.Args[2]
	message := os.Args[3]

	discordBot, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println(err)
	}

	err = discordBot.Open()
	if err != nil {
		fmt.Println("Error on Open:", err)
	}

	_, err = discordBot.ChannelMessageSend(channelID, message)

	if err != nil {
		fmt.Println("Error on send:", err)
	}
}
