package pkg

import (
	"time"
)

type Config struct {
	Url           string `json:"url"`
	ApiKey        string `json:"api_key"`
	ActiveProgram string `json:"active_program"`
}

type Program struct {
	ID         int    `json:"id"`
	PlatformID int    `json:"platform_id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	VDP        bool   `json:"vdp"`
	Favourite  bool   `json:"favourite"`
	Tag        string `json:"tag"`
	URL        string `json:"url"`
	Type       string `json:"type"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type Scope struct {
	ID        int64     `json:"id"`
	ProgramID int64     `json:"program_id"`
	Scope     string    `json:"scope"`
	ScopeType string    `json:"scope_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Subdomain struct {
	ID            int64     `json:"id"`
	ProgramID     int64     `json:"program_id"`
	Subdomain     string    `json:"subdomain"`
	Tag           string    `json:"tag"`
	Ip            *int64    `json:"ip"`
	Title         string    `json:"title"`
	BodyHash      string    `json:"body_hash"`
	StatusCode    int32     `json:"status_code"`
	ContentLength int32     `json:"content_length"`
	Screenshot    string    `json:"screenshot"`
	Favourite     bool      `json:"favourite"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TechnologieSubdomain struct {
	ID                int64     `json:"id"`
	TechnologyVersion int64     `json:"technology_version"`
	SubdomainID       int64     `json:"subdomain_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type TechnologieVersion struct {
	ID           int64     `json:"id"`
	TechnologyID int64     `json:"technology_id"`
	Version      string    `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Technology struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	ID        int64     `json:"id"`
	ProgramID int64     `json:"program_id"`
	Ip        string    `json:"ip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Port struct {
	ID        int64     `json:"id"`
	IpID      int64     `json:"ip_id"`
	Port      int32     `json:"port"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WappaGoTechnology struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Cpe     string `json:"cpe"`
}

type WappaGoInfo struct {
	StatusCode    int                 `json:"status_code"`
	Ports         []string            `json:"ports"`
	Path          string              `json:"path"`
	Location      string              `json:"location"`
	Title         string              `json:"title"`
	Scheme        string              `json:"scheme"`
	Data          string              `json:"data"`
	ResponseTime  int                 `json:"response_time"`
	Screenshot    string              `json:"screenshot_name"`
	Technologies  []WappaGoTechnology `json:"technologies"`
	ContentLength int                 `json:"content_length"`
	ContentType   string              `json:"content_type"`
	IP            string              `json:"ip"`
	CertVHost     []string            `json:"certvhost"`
}

type WappaGo struct {
	Url   string      `json:"url"`
	Infos WappaGoInfo `json:"infos"`
}
