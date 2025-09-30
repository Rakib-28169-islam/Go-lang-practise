package blog

import (
	"encoding/json"
	// "io/ioutil"
	"sync"
	"time"

)

type Post struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    Author    string `json:"author"`
    Published bool   `json:"published"`
    CreatedAt string `json:"created_at"`
}
type Blog struct{
	posts map[int]Post
	nextID int
	mu sync.RWMutex
}

func newBlog() *Blog{
	return &Blog{
		posts: make(map[int]Post),
		nextID: 1,
	}
}
func (b *Blog) AddPost(title, content,author string,published bool){
	b.mu.Lock()
	defer b.mu.Unlock()
	post := Post{
		ID: b.nextID,
		Title: title,
		Content: content,
		Author: author,
		Published: published,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),

    }
	b.posts[b.nextID] = post
	b.nextID++
}
func (b *Blog) GetPost(id int) (string, bool){
	b.mu.RLock()
	defer b.mu.RUnlock()
	post, exists := b.posts[id]
	if !exists{
		return "",false
	}
	postJSON,err := json.Marshal(post)
	if err != nil{
		return "failed to create json",false
	}
	return string(postJSON), exists
}
func (b *Blog) GetAllPosts() string {
    b.mu.RLock()//
    defer b.mu.RUnlock()//
    
    posts := make([]Post, 0, len(b.posts))//
    for _, post := range b.posts {
        posts = append(posts, post)
    }
    
    jsonStr, _ := json.Marshal(posts)

    return string(jsonStr)
}