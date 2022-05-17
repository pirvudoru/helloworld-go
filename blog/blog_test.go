package blogposts_test

import (
	"errors"
	blogposts "github.com/pirvudoru/blogposts"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
	"testing/fstest"
)

var MyFSError = errors.New("oh no, i always fail")

type StubFailingFS struct {
}

func (s StubFailingFS) Open(_ string) (fs.File, error) {
	return nil, MyFSError
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("when failing to read the FS", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		assert.Equal(t, MyFSError, err)
	})

	t.Run("when", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
		)

		myFS := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogposts.NewPostsFromFS(myFS)

		assert.Equal(t, posts[0], blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})
	})
}
