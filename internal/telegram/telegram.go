package telegram

import (
	"fmt"
	"log"
	"movie_suggestion/constants"
	"movie_suggestion/internal/movie-api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotGetMovie interface {
	GetMovie(genre string) (m movie.Movie, err error)
}

type Bot struct {
	token string
	bot   *tgbotapi.BotAPI
	movie BotGetMovie
}

func New(movie BotGetMovie, token string) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	return &Bot{
		token: token,
		bot:   bot,
		movie: movie,
	}
}

func (b *Bot) Start() {

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates, err := b.bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !checkValidMessageText(update) && update.Message.Text != "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "There is no such command")
			if _, err := b.bot.Send(msg); err != nil {
				log.Fatal(err)
			}
			continue
		}

		movie_answer_message, err := b.movie.GetMovie(update.Message.Text)
		if err != nil {
			log.Fatal(err)
		}

		message_text := fmt.Sprintf("🔥Title: %s\n\n📊IMDb Rating: %s\n\n🎭Genres: %s\n\n📒Plot: %s\n\n🔗Link: %s", movie_answer_message.Title, movie_answer_message.Rating, movie_answer_message.Genres, movie_answer_message.Plot, "https://www.imdb.com/title/"+movie_answer_message.MovieID)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message_text)

		switch update.Message.Command() {
		case "start":
			keyboardSetMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Choose genre from the keyboard below")
			keyboardSetMsg.ReplyMarkup = constants.NumericKeyboard
			if _, err := b.bot.Send(keyboardSetMsg); err != nil {
				log.Fatal(err)
			}
			continue
		}

		if _, err := b.bot.Send(msg); err != nil {
			log.Fatal(err)
		}
	}

}

func checkValidMessageText(update tgbotapi.Update) bool {
	for _, g := range constants.Genres {
		if update.Message.Text == g {
			return true
		}
	}
	return false
}
