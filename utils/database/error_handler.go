package database

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
)

func ErrorHandler(err error) error {
	// print log in case some errors is not handled and returned as "unknown"
	fmt.Printf("database ErrorHandler: %v, type: %T\n", err, err.Error)

	// try to assert error as PgError
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		fmt.Printf("PostgreSQL error: %v\n", pgErr)

		// handle different error codes here
		switch pgErr.Code {
		case "23505":
			return errors.New("already_exists")
		// all other codes
		default:
			return errors.New("unknown_error")
		}
	} else {
		fmt.Println("error is not a *pgconn.PgError")
	}

	return errors.New("unknown_error")
}
