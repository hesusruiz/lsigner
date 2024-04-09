package certstore

import "errors"

// This will hopefully give a compiler error that will hint at the fact that
// this package isn't designed to work on Linux.
func init() {

}

// Implement this function, just to silence other compiler errors.
func openStore() (Store, error) {
	return nil, errors.New("certstore only works on macOS and Windows")
}

func UseUserStore() { panic("certstore only works on macOS and Windows") }
