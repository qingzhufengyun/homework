package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func get_record()(err error) {
	err = sql.ErrNoRows
	fmt.Errorf("access sql error %w", err)

	return
}

func main() {
	var err error
	err = get_record()

	if errors.Is(err, sql.ErrNoRows) {
       fmt.Println("err is sql.ErrNoRows")
	}else {
		fmt.Println("err isn't sql.ErrNoRows")
	}
}

