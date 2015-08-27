package main

import (
	"errors"
	"sync"
)

var errNotFound = errors.New("Article not found")

// Article is content posted by a user
type Article struct {
	ID   int    `json:"id"`
	User string `json:"user"`
	Body string `json:"body"`
}

// DB holds our Articles
type DB struct {
	sync.RWMutex
	articles []Article
}

// FindOne finds an article by its ID
func (db *DB) FindOne(id int) (Article, error) {
	db.RLock()
	defer db.RUnlock()

	for _, p := range db.articles {
		if p.ID == id {
			return p, nil
		}
	}

	return Article{}, errNotFound
}

// FindAll returns all Articles
func (db *DB) FindAll() []Article {
	db.RLock()
	defer db.RUnlock()

	return db.articles
}

// Insert adds an Article to the DB. It sets the ID field and returns the modified
// Article.
func (db *DB) Insert(p Article) Article {
	db.Lock()
	defer db.Unlock()

	id := 0

	for _, article := range db.articles {
		if article.ID > id {
			id = article.ID
		}
	}
	id++

	p.ID = id

	db.articles = append(db.articles, p)

	return p
}

// Update finds an Article in the DB by ID and replaces it
func (db *DB) Update(p Article) {
	db.Lock()
	defer db.Unlock()

	for i, article := range db.articles {
		if article.ID == p.ID {
			db.articles[i] = p
			return
		}
	}
}

// Delete removes an Article from the collection by ID
func (db *DB) Delete(id int) {
	db.Lock()
	defer db.Unlock()

	for i, article := range db.articles {
		if article.ID == id {
			db.articles = append(db.articles[:i], db.articles[i+1:]...)
			return
		}
	}
}
