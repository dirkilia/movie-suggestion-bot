package movie

type MoviesList struct {
	Movies_list []Movie `json:"results"`
}

type Movie struct {
	MovieID string `json:"id"`
	Title   string `json:"title"`
	Plot    string `json:"plot"`
	Image   string `json:"image"`
	Genres  string `json:"genres"`
	Rating  string `json:"imDbRating"`
}
