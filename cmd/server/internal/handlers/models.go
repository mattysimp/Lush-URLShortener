package handlers

import (
	"hash/crc32"
	"strconv"
)

// URL holds LongURL, ShorURL and URLCode for a URL
type URL struct {
	LongURL  string `json:"url"`
	ShortURL string `json:"short_url"`
	URLCode  string `json:"url_code" bson:"_id"`
}

type Config struct {
	BaseURL  string   `json:"base_url"`
	Database database `json:"database"`
}

type database struct {
	Host       string `json:"host"`
	DBName     string `json:"db_name"`
	Collection string `json:"collection"`
}

func (url *URL) GenerateShortURL(config *Config, amendment string) {
	// Generate URL
	code := crc32.ChecksumIEEE([]byte(url.LongURL))
	url.URLCode = strconv.FormatUint(uint64(code), 36) + amendment
	url.ShortURL = config.BaseURL + url.URLCode
}
