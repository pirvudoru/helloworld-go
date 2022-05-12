package dictionary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("other")

		assert.Equal(t, "", got)
		assert.Equal(t, ErrNotFound, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)
		assert.Nil(t, err)

		got, err := dictionary.Search(word)
		assert.Equal(t, definition, got)
		assert.Nil(t, err)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		assert.Equal(t, ErrWordExists, err)

		value, err := dictionary.Search(word)
		assert.Nil(t, err)
		assert.Equal(t, definition, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		originalDefinition := "this is just a test"
		dictionary := Dictionary{word: originalDefinition}
		newDefinition := "this is just a test"
		err := dictionary.Update(word, newDefinition)
		assert.Nil(t, err)

		value, err := dictionary.Search(word)
		assert.Nil(t, err)
		assert.Equal(t, newDefinition, value)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Update(word, definition)
		assert.Equal(t, ErrWordDoesNotExist, err)
	})
}
func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
