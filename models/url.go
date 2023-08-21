package models

import (
	"time"

	"github.com/HectorMenezes/url-shortener-go/utils"

	"github.com/jinzhu/gorm"
)

// Url represents a table of database and a payload.
type Url struct {
    ID        string    `json:"id" example:"1534637"`
    Url       string    `json:"url" validate:"required,url" example:"chess.com"`
    CreatedAt time.Time `json:"createdAt" example:"2023-08-20T23:41:35Z"`
}


// BeforeCreate sets values missing, both of them need to be
// generated right beforte creating.
func (url *Url) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("ID", utils.Hash(7, url.Url))
	return nil
}

type UrlRequest struct {
    // Url in the format `https://mysite.com`
    Url       string    `json:"url" validate:"required,url" example:"https://chess.com"`
} // @name UrlInformation

// UrlResponse represents an output payload
type UrlResponse struct {
    //Link created
    Link        string `json:"link" binding:"required,url" example:"PREFIX/1534637"`
    //Original URL
    OriginalUrl string `json:"originalUrl" binding:"rquired,url" example:"https://chess.com"`
    //When it was created
    CreatedAt time.Time `json:"createdAt" example:"2023-08-20T23:41:35Z"`
} // @name UrlResponse
