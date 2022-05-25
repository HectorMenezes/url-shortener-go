package models

import (
	"time"

	"github.com/HectorMenezes/url-shortener-go/utils"

	"github.com/jinzhu/gorm"
)

// Url represents a table of database and a payload.
type Url struct {
	ID        string    `json:"id"`
	Url       string    `json:"url" validate:"required,url"`
	CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate sets values missing, both of them need to be
// generated right beforte creating.
func (url *Url) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("ID", utils.Hash(7, url.Url))
	return nil
}

// UlrResponsePost represents an output payload
type UlrResponsePost struct {
	Link        string `json:"link" binding:"required,url"`
	OriginalUrl string `json:"original_url" binding:"rquired,url"`
}
