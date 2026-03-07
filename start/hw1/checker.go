package main

import (
	"fmt"
)

const maxTimeOut = 120

func checkCommand(a *Args) error {
	if len(a.Urls) < 1 {
		return fmt.Errorf("given 0 urls, required 1 at least")
	}

	if a.Time < 0 {
		return fmt.Errorf("timeout value can't be negative: %d", a.Time)
	}

	if a.Time > maxTimeOut {
		return fmt.Errorf("timeout value must not be more 120s")
	}

	return nil
}
