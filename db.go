package main

import (
	"errors"
	"sync"
)

var errNotFound = errors.New("Post not found")

// Post is a message by a user
type Post struct {
	ID      int    `json:"id"`
	User    string `json:"user"`
	Message string `json:"message"`
}

// DB holds our Posts
type DB struct {
	sync.RWMutex
	posts []Post
}

// FindOne finds a post by its ID
func (db *DB) FindOne(id int) (Post, error) {
	db.RLock()
	defer db.RUnlock()

	for _, p := range db.posts {
		if p.ID == id {
			return p, nil
		}
	}

	return Post{}, errNotFound
}

// FindAll returns all Posts
func (db *DB) FindAll() []Post {
	db.RLock()
	defer db.RUnlock()

	return db.posts
}

// Insert adds a Post to the DB. It sets the ID field and returns the modified
// Post.
func (db *DB) Insert(p Post) Post {
	db.Lock()
	defer db.Unlock()

	id := 0

	for _, post := range db.posts {
		if post.ID > id {
			id = post.ID
		}
	}
	id++

	p.ID = id

	db.posts = append(db.posts, p)

	return p
}

// Update finds a Post in the DB by ID and replaces it
func (db *DB) Update(p Post) {
	db.Lock()
	defer db.Unlock()

	for i, post := range db.posts {
		if post.ID == p.ID {
			db.posts[i] = p
			return
		}
	}
}

// Delete removes a Post from the collection by ID
func (db *DB) Delete(id int) {
	db.Lock()
	defer db.Unlock()

	for i, post := range db.posts {
		if post.ID == id {
			db.posts = append(db.posts[:i], db.posts[i+1:]...)
			return
		}
	}
}
