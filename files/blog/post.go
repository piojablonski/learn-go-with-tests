package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleToken       = "Title: "
	descriptionToken = "Description: "
	tagsToken        = "Tags: "
	bodyToken        = "---"
)

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)

	readMetaLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	readTags := func() []string {
		tags := strings.Split(readMetaLine(tagsToken), ",")
		var res []string
		for _, tag := range tags {
			res = append(res, strings.TrimSpace(tag))
		}
		return res
	}

	return Post{
		Title:       readMetaLine(titleToken),
		Description: readMetaLine(descriptionToken),
		Tags:        readTags(),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {

	var buf bytes.Buffer

	// ignore prefix
	scanner.Scan()
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
