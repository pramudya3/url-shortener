package domain

import "time"

type URL struct {
	URLRaw    string    `json:"url_raw"`
	URLShort  string    `json:"url_short"`
	CreatedAt time.Time `json:"created_at"`
}
