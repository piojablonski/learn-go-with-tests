package main

import (
	"files/blog"
	"fmt"
	"os"
)

func main() {
	posts, err := blog.NewPostsFromFS(os.DirFS("posts"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", posts)
}
