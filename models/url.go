package models

import (
	"time"

	"github.com/HectorMenezes/url-shortener-go/utils"

	"github.com/jinzhu/gorm"
)

type Url struct {
	ID        string    `json:"id"`
	Url       string    `json:"url" validate:"required,url"`
	CreatedAt time.Time `json:"created_at"`
}

func (url *Url) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("ID", utils.Hash(7, url.Url))

	return nil
}

type UlrResponsePost struct {
	Link        string `json:"link" binding:"required,url"`
	OriginalUrl string `json:"original_url" binding:"rquired,url"`
}
