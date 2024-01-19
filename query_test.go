package pgglue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Select(t *testing.T) {
	t.Run("no columns", func(t *testing.T) {
		q := Select("books")
		assert.NotNil(t, q)
		assert.Equal(t, "select * from books;", q.S())
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

func Test_Update(t *testing.T) {
	q := Update("books", []string{"author", "publisher", "publish_date"})
	assert.NotNil(t, q)
	assert.Equal(t, "update books set ( author = $1, publisher = $2, publish_date = $3 );", q.S())
}

func Test_Delete(t *testing.T) {
	q := Delete("books")
	assert.NotNil(t, q)
	assert.Equal(t, "delete from books;", q.S())
}

func Test_Where(t *testing.T) {
	q := Update("books", []string{"author", "publisher", "publish_date"})
	assert.NotNil(t, q)
	assert.Equal(t, "update books set ( author = $1, publisher = $2, publish_date = $3 )", q.statement)
	q = q.Where("author", "=")
	assert.Equal(t, "update books set ( author = $1, publisher = $2, publish_date = $3 ) where author = $4;", q.S())
}

func Test_Returning(t *testing.T) {
	columns := []string{"author", "publisher", "publish_date"}
	q := Update("books", columns).
		Where("author", "=").
		Returning(columns)
	assert.NotNil(t, q)
	statement := "update books set ( author = $1, publisher = $2, publish_date = $3 ) where author = $4 returning ( author, publisher, publish_date );"
	assert.Equal(t, statement, q.S())
}

func Test_Limit(t *testing.T) {
	q := Select("books").Limit(50)
	assert.NotNil(t, q)
	assert.Equal(t, "select * from books limit 50;", q.S())
}

func Test_Join(t *testing.T) {
	q := Select("books").Join("reviews", "books.id = reviews.book_id")
	assert.NotNil(t, q)
	assert.Equal(t, "select * from books join reviews on books.id = reviews.book_id;", q.S())
}
