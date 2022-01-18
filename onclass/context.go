package main

import "net/http"

// why not interface
type Context struct {
	W http.ResponseWriter
	R *http.Request
}
