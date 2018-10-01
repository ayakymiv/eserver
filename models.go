package main

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	ID uint `json:”id”`
	Name string `json:"name"`
	Event string `json:"event"`
	Handlers Handler
	data interface{}
}

type Handler struct{
	gorm.Model
	ID uint `json:”id”`
	Handler string

}
