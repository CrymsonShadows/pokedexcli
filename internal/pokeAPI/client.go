package pokeapi

import (
	"net/http"
	"time"

	"github.com/CrymsonShadows/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(20 * time.Minute),
	}
}
