package constants

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var NumericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Comedy"),
		tgbotapi.NewKeyboardButton("Action"),
		tgbotapi.NewKeyboardButton("History"),
		tgbotapi.NewKeyboardButton("Horror"),
		tgbotapi.NewKeyboardButton("Thriller"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Adventure"),
		tgbotapi.NewKeyboardButton("Crime"),
		tgbotapi.NewKeyboardButton("Fantasy"),
		tgbotapi.NewKeyboardButton("Drama"),
		tgbotapi.NewKeyboardButton("Documentary"),
	),
)

var Genres []string = []string{"Comedy", "Action", "History", "Horror", "Thriller", "Adventure", "Crime", "Fantasy", "Drama", "Documentary"}
