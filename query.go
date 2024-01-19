package pgglue

import "strconv"

// Query is a struct that contains a psql statement
type Query struct {
	statement string
}

// S returns the query statement, built and ready for execution.
func (q *Query) S() string {
	return q.statement + ";"
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
		statement: statement,
	}
}

// func Update() {}

// func Delete() {}

// func Returning() {}

// func Limit() {}
