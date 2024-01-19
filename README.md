# pgglue

A very simple query builder for postgres

## Features

### Select

Without specified columns

```go
q := Select("books")
```

```sql
select *  from books;
```

With specified columns

```go
q := Select("books", []string{"author", "publisher", "publish_date"})
```

```sql
select author, publisher, publish_date from books;
```

### Limit

```go
q := Select("books").Limit(50).S()
```

```sql
select *  from books limit 50;
```

### Insert

```go
q := Insert("books", []string{"author", "publisher", "publish_date"}).S()
```

```sql
insert into books (author, publisher, publish_date) values ($1, $2, $3);
```

### Update

```go
q := Update("books", []string{"author", "publisher", "publish_date"}).S()
```

```sql
update books set ( author = $1, publisher = $2, publish_date = $3 );
```

### Where

```go
q := Update("books", []string{"author", "publisher", "publish_date"}).Where("author", "=").S()
```

```sql
update books set ( author = $1, publisher = $2, publish_date = $3 ) where author = $4;
```

### Returning

```go
	columns := []string{"author", "publisher", "publish_date"}
	q := Update("books", columns).
		Where("author", "=").
		Returning(columns)
```

```sql
update books set ( author = $1, publisher = $2, publish_date = $3 ) where author = $4 returning ( author, publisher, publish_date );
```

### Delete

```go
q := Delete("books")
```

```sql
delete from books;
```

## Example Usage

```go
type Book struct {
	Author      string
	Publisher   string
	PublishDate time.Time
}

type GetArgs struct {
	AuthorID       string
	IncludeReviews bool
	IncludeRatings bool
}

func GetByAuthor(ctx context.Context, args GetArgs) ([]Book, error) {
	q := pgglue.Select("books")
	if args.IncludeReviews {
		q.Join("reviews", "books.id = reviews.book_id")
	}

	if args.IncludeRatings {
		q.Join("ratings", "books.id = reviews.book_id")
	}

	rows := yourDriverOfChoice.Query(
		ctx,
		q.S(),
		args.AuthorID,
	)

	// Scan and return your results...
}
```
