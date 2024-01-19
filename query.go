package pgglue

import (
	"fmt"
	"strconv"
)

// Query is a struct that contains a psql statement
type Query struct {
	argNum    int
	statement string
}

// S returns the query statement, built and ready for execution.
func (q *Query) S() string {
	return q.statement + ";"
}

// Where adds a clause to an existing query's statement. The arguments
// required are the column being evaluated and the type of expression to use (>, <, =)
func (q *Query) Where(column, expression string) *Query {
	q.argNum++
	q.statement += fmt.Sprintf(" where %s %s $%s", column, expression, strconv.Itoa(q.argNum))
	return q
}

// Returning adds a returning clause to the query statement from a list of provided columns
func (q *Query) Returning(columns []string) *Query {
	q.statement += " returning ( "
	for i, c := range columns {
		q.statement += c
		if i != len(columns)-1 {
			q.statement += ", "
		}
	}
	q.statement += " )"
	return q
}

// Limit adds a limit to a Query's statement
func (q *Query) Limit(limit int) *Query {
	q.statement += " limit " + strconv.Itoa(limit)
	return q
}

// Select takes a table and column names for a select query
func Select(table string, columns ...[]string) *Query {
	if table == "" {
		return nil
	}

	var statement string
	if len(columns) == 0 {
		statement = "select * "
	}

	if statement == "" {
		statement = "select "
		cols := columns[0]
		for i, c := range cols {
			statement += c
			if i != len(cols)-1 {
				statement += ", "
			}
		}
	}

	statement += " from " + table
	return &Query{statement: statement}
}

// Insert takes a table and arguments for an insert query and crafts the psql statement
func Insert(table string, args []string) *Query {
	if table == "" || len(args) == 0 {
		return nil
	}
	statement := "insert into " + table + " ("
	values := "values ("
	for i, a := range args {
		statement += a
		values += "$" + strconv.Itoa(i+1)
		if i != len(args)-1 {
			statement += ", "
			values += ", "
		}
	}
	statement += ") " + values + ")"
	return &Query{
		argNum:    len(args),
		statement: statement,
	}
}

// Update takes a table and a list of columns meant to be updated.
// The query returned should be scoped with .Where()
func Update(table string, args []string) *Query {
	if table == "" || len(args) == 0 {
		return nil
	}
	statement := "update " + table + " set ( "
	for i, a := range args {
		statement += a + " = $" + strconv.Itoa(i+1)
		if i != len(args)-1 {
			statement += ", "
		}
	}
	statement += " )"
	return &Query{
		argNum:    len(args),
		statement: statement,
	}
}

// Delete takes a table name from which records will be deleted.
// The query returned should be scoped with .Where()
func Delete(table string) *Query {
	if table == "" {
		return nil
	}

	return &Query{
		statement: "delete from " + table,
	}
}
