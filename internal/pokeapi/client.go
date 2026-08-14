package pokeapi

import (
	"time"
	"github.com/mirmeep/pokedex-cli/internal/pokecache"
)

// Client -
type Client struct {
	cache      *pokecache.Cache 
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
	}
}
