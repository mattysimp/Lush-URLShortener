package handlers

import (
	"hash/crc32"
	"strconv"
)

// URL holds LongURL, ShorURL and URLCode for a URL
type URL struct {
	LongURL  string `json:"url"`
	ShortURL string `json:"short_url"`
	URLCode  uint32 `json:"url_code"`
}

type Config struct {
	BaseURL string `json:"base_url"`
}

func (url *URL) GenerateShortURL(config *Config) {
	// Generate URL
	code := crc32.ChecksumIEEE([]byte(url.LongURL))
	url.URLCode = code
	url.ShortURL = config.BaseURL + strconv.FormatUint(uint64(code), 62)
}
