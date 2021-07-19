package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"botFFMPEG/actions"
	"botFFMPEG/checkers"
	"botFFMPEG/commands"

	"github.com/bwmarrin/discordgo"
)

// Переменные, используемые для параметров командной строки
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Создание сеанса с использованием токена от бота
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	// Инициализация папку для работы с ffmpeg
	if !checkers.CheckFile("videos") {
		err = os.Mkdir("videos", 0777)
		if err != nil {
			panic(err)
		}
	}
	// колбеки для функций MessageCreate
	dg.AddHandler(commands.EncHandlerWebm)
	dg.AddHandler(commands.EncHandlerHevc)
	dg.AddHandler(actions.WaitMessage)
	// Обработка сообщений только с серверов
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	// Открытие соединение с дискордом
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	// Бот запущен
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
