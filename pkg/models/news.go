package models

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

// News type
type News struct {
	ID        string
	Title     string
	Image     string
	Detail    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var newsStorage []*News

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

// NewsList list of news
func NewsList() []*News {
	return newsStorage
}

// CreateNews create new news
func CreateNews(news *News) {
	news.ID = generateID()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt
	newsStorage = append(newsStorage, news)
}

// GetNews get news by id
func GetNews(id string) *News {
	for _, news := range newsStorage {
		if news.ID == id {
			return news
		}
	}
	return nil
}

// DeleteNews delete a news
func DeleteNews(id string) {
	for i, news := range newsStorage {
		if news.ID == id {
			newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
			return
		}
	}
}
