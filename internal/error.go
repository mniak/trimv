package internal

import (
	"fmt"
	"os"
)

func Handle(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
