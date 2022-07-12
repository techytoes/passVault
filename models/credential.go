package models

import "time"

type Credential struct {
	Username    []byte    `json:"username"`
	Password    []byte    `json:"password"`
	App         string    `json:"app"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	LastUsed    time.Time `json:"lastUsed"`
	Version     int16     `json:"version"`
}
