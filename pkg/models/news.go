package models

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
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

var newsStorage []News
var mutxNews sync.RWMutex

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

// NewsList list of news
func NewsList() []*News {
	mutxNews.RLock()
	defer mutxNews.RUnlock()
	r := make([]*News, len(newsStorage))

	for i := range newsStorage {
		n := newsStorage[i]
		r[i] = &n
	}
	return r
}

// CreateNews create new news
func CreateNews(news News) {
	news.ID = generateID()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt

	mutxNews.Lock()
	defer mutxNews.Unlock()
	newsStorage = append(newsStorage, news)
}

// GetNews get news by id
func GetNews(id string) *News {
	mutxNews.RLock()
	defer mutxNews.RUnlock()

	for _, news := range newsStorage {
		if news.ID == id {
			n := news
			return &n
		}
	}
	return nil
}

// DeleteNews delete a news
func DeleteNews(id string) {
	mutxNews.Lock()
	defer mutxNews.Unlock()
	for i, news := range newsStorage {
		if news.ID == id {
			newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
			return
		}
	}
}
