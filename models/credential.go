package models

import "time"

type Credential struct {
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Password    []byte    `json:"password"`
	App         string    `json:"app"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	LastUsed    time.Time `json:"lastUsed"`
}
