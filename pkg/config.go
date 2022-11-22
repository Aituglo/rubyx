package pkg

import "time"

type Config struct {
	Url             string `json:"url"`
	ApiKey          string `json:"api_key"`
	ActiveProgram   string `json:"active_program"`
	ActiveProgramID int    `json:"active_program_id"`
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
