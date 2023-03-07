package main

import (
	"log"
	"movie_suggestion/internal/movie-api"
	"movie_suggestion/internal/telegram"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print(err)
	}

	token, ok := os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		log.Println("token not found")
	}
	apikey, ok := os.LookupEnv("IMDB_API_KEY")
	if !ok {
		log.Println("api key not found")
	}

	movie_answer_message := movie.New(apikey)
	movieBot := telegram.New(movie_answer_message, token)
	movieBot.Start()
}
