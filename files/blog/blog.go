package blog

import (
	"fmt"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	files, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, fmt.Errorf("new posts from fs: %w", err)
	}
	var posts []Post
	for _, f := range files {
		post, err := openPostFile(fileSystem, f.Name())
		if err != nil {
			return nil, fmt.Errorf("creating posts from fs %w", err)
		}
		posts = append(posts, post)
	}
	return posts, nil

}

func openPostFile(fileSystem fs.FS, fileName string) (Post, error) {
	file, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, fmt.Errorf("open post file: %w", err)
	}
	defer file.Close()
	return newPost(file)
}
