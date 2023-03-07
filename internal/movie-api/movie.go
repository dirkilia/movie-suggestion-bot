package movie

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type getMovie struct {
	apikey string
}

func New(apiKey string) *getMovie {
	return &getMovie{
		apikey: apiKey,
	}
}

func (g *getMovie) GetMovie(genre string) (m Movie, err error) {
	var data MoviesList

	var request string = fmt.Sprintf("https://imdb-api.com/API/AdvancedSearch/" +
		g.apikey +
		"?genres=" +
		genre +
		"&title_type=feature,tv_movie,documentary,short,tv_short&count=250")

	resp, err := http.Get(request)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	if len(data.Movies_list) > 0 {
		m = getSingleMovie(data.Movies_list)
	} else {
		log.Fatal(err)
	}
	return m, nil
}

func getSingleMovie(list []Movie) Movie {
	rand.Seed(time.Now().UnixNano())
	movie_number := rand.Intn(len(list))

	return list[movie_number]
}
