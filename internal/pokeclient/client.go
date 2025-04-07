package pokeclient

import (
	"net/http"
	"time"

	"github.com/petomackay/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache *pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
