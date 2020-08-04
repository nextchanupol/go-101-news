package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// News type
type News struct {
	ID        bson.ObjectId `bson:"_id"`
	Title     string
	Image     string
	Detail    string
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

var newsStorage []News
var mutxNews sync.RWMutex

func generateID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf)
}

// NewsList list of news
func NewsList() ([]*News, error) {

	s := mongoSession.Copy()
	defer s.Close()
	var news []*News
	err := s.DB(database).C("news").Find(nil).All(&news)
	if err != nil {
		return nil, err
	}
	return news, nil

	// mutxNews.RLock()
	// defer mutxNews.RUnlock()
	// r := make([]*News, len(newsStorage))

	// for i := range newsStorage {
	// 	n := newsStorage[i]
	// 	r[i] = &n
	// }
	// return r
}

// CreateNews create new news
func CreateNews(news News) error {
	news.ID = bson.NewObjectId()
	news.CreatedAt = time.Now()
	news.UpdatedAt = news.CreatedAt

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").Insert(&news)
	if err != nil {
		return err
	}
	return nil

	// mutxNews.Lock()
	// defer mutxNews.Unlock()
	// newsStorage = append(newsStorage, news)
}

// GetNews get news by id
func GetNews(id string) (*News, error) {

	if !bson.IsObjectIdHex(id) {
		return nil, fmt.Errorf("invalid id")

	}
	objectID := bson.ObjectIdHex(id)

	s := mongoSession.Copy()
	defer s.Close()
	var n News
	err := s.DB(database).C("news").FindId(objectID).One(&n)

	if err != nil {
		return nil, err
	}

	return &n, nil

	// mutxNews.RLock()
	// defer mutxNews.RUnlock()

	// for _, news := range newsStorage {
	// 	if news.ID == id {
	// 		n := news
	// 		return &n
	// 	}
	// }
	// return nil
}

// DeleteNews delete a news
func DeleteNews(id string) error {

	objectID := bson.ObjectId(id)

	if !objectID.Valid() {
		return fmt.Errorf("invalid id")
	}

	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("news").RemoveId(objectID)

	if err != nil {
		return err
	}
	return nil

	// mutxNews.Lock()
	// defer mutxNews.Unlock()
	// for i, news := range newsStorage {
	// 	if news.ID == id {
	// 		newsStorage = append(newsStorage[:i], newsStorage[i+1:]...)
	// 		return
	// 	}
	// }
}
