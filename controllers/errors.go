package controllers

import "errors"

var (
	InvalidUUID = errors.New("invalid uuid")
	AccessError = errors.New("access uuid")
)
