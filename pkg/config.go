package pkg

import "time"

type Config struct {
	Url           string `json:"url"`
	ApiKey        string `json:"api_key"`
	ActiveProgram string `json:"active_program"`
}

type Program struct {
	ID         int64     `json:"id"`
	PlatformID int64     `json:"platform_id"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	Vdp        bool      `json:"vdp"`
	Url        string    `json:"url"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Subdomain struct {
	ID            int64     `json:"id"`
	ProgramID     int64     `json:"program_id"`
	Url           string    `json:"url"`
	Title         string    `json:"title"`
	BodyHash      string    `json:"body_hash"`
	StatusCode    int32     `json:"status_code"`
	Technologies  string    `json:"technologies"`
	ContentLength int32     `json:"content_length"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Url struct {
	ID            int64     `json:"id"`
	SubdomainID   int64     `json:"subdomain_id"`
	Url           string    `json:"url"`
	Title         string    `json:"title"`
	BodyHash      string    `json:"body_hash"`
	StatusCode    int32     `json:"status_code"`
	Technologies  string    `json:"technologies"`
	ContentLength int32     `json:"content_length"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Ip struct {
	ID          int64     `json:"id"`
	ProgramID   int64     `json:"program_id"`
	SubdomainID int64     `json:"subdomain_id"`
	Ip          string    `json:"ip"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
