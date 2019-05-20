package handler

import (
	"go-echo-mongo/go-echo-mongo/mware"
)

type (

	Handler struct {
		DB mware.DB
})


const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
