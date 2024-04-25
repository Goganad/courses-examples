package main

import (
	"errors"
	"fmt"
)

func main() {
	err := runQuery("connect")
	if err != nil {
		var connErr *ConnectionError
		if errors.As(err, &connErr) {
			fmt.Printf("Query error: %s, %+v", connErr.Msg, err)
		}

		var queryErr *QueryError
		if errors.As(err, &queryErr) {
			fmt.Printf("%+v", err)
		}
	}

	err = runQuery("select")
	if err != nil {
		var queryErr *QueryError
		if errors.As(err, &queryErr) {
			fmt.Println("Query error:", queryErr)
		}
	}
}

func multiple() (int, int, error) {
	return 0, 0, nil
}

type ConnectionError struct {
	Msg string
}

func (e *ConnectionError) Error() string {
	return e.Msg
}

type QueryError struct {
	Query string
	Msg   string
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("error in query '%s': %s", e.Query, e.Msg)
}

func runQuery(query string) error {
	if query == "connect" {
		return &ConnectionError{"failed to connect to database"}
	} else if query == "select" {
		return &QueryError{Query: query, Msg: "syntax error"}
	}
	return nil
}
