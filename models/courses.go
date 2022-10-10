package models

type Course struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       string `json:"level"`
}

var Courses []Course
