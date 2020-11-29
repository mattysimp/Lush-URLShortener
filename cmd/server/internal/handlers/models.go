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

// Config holds information provided by config.json
type Config struct {
	BaseURL  string   `json:"base_url"`
	Database database `json:"database"`
}

// Database holds the database information provided by config.json
type database struct {
	Host       string `json:"host"`
	DBName     string `json:"db_name"`
	Collection string `json:"collection"`
}

// GenerateShortURL hashes LongURL to creat URLCode and ShortURL
func (url *URL) GenerateShortURL(config *Config, amendment string) {
	code := crc32.ChecksumIEEE([]byte(url.LongURL))
	url.URLCode = strconv.FormatUint(uint64(code), 36) + amendment
	url.ShortURL = config.BaseURL + url.URLCode
}
