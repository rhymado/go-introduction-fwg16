package models

type Post struct {
	Title   string
	Content string
	Owner   User
}

func CreatePost(title, content string, owner User) Post {
	return Post{
		Title:   title,
		Content: content,
		Owner:   owner,
	}
}

func (p *Post) EditPost(title, content string) {
	p.Title = title
	p.Content = content
}
