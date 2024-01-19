package pgglue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Select(t *testing.T) {
	t.Run("no columns", func(t *testing.T) {
		q := Select("books")
		assert.NotNil(t, q)
		assert.Equal(t, "select *  from books;", q.S())
	})

	t.Run("with columns", func(t *testing.T) {
		q := Select("books", []string{"author", "publisher", "publish_date"})
		assert.NotNil(t, q)
		assert.Equal(t, "select author, publisher, publish_date from books;", q.S())
	})
}
func Test_Insert(t *testing.T) {
	q := Insert("books", []string{"author", "publisher", "publish_date"})
	assert.NotNil(t, q)
	assert.Equal(t, "insert into books (author, publisher, publish_date) values ($1, $2, $3);", q.S())
}
