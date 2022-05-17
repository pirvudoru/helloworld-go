package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}

	defer func() {
		_ = postFile.Close()
	}()
	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
	tagSeparator         = ", "
)

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	return Post{
		Title:       readMetaLine(scanner, titleSeparator),
		Description: readMetaLine(scanner, descriptionSeparator),
		Tags:        readTags(scanner),
		Body:        readBody(scanner),
	}, nil
}

func readMetaLine(scanner *bufio.Scanner, tagName string) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), tagName)
}

func readTags(scanner *bufio.Scanner) []string {
	return strings.Split(readMetaLine(scanner, tagsSeparator), tagSeparator)
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line
	buf := bytes.Buffer{}
	for scanner.Scan() {
		_, _ = fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
