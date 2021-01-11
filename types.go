package main

import "time"

type softwareDescription struct {
	ID       int `json:"-"`
	Name     string
	Internal string `json:"-"`
}

type subscription struct {
	Name              string
	LastLoginDate     time.Time
	DaysTillEnd       int
	Key               string    `json:"-"`
	AvailableWarezIDs []int     `json:"-"`
	Created           time.Time `json:"-"`
	SubDays           int       `json:"-"`
}
