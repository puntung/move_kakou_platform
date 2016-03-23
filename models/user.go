package models

import (
	"time"
)

type Users struct {
	Id           int       `orm:"column(id)"`
	Username     string    `orm:"column(username)";size(32)`
	Password     string    `orm:column(password);size(255)`
	Scope        string    `orm:"column(scpoe);size(255)"`
	DateCreated  time.Time `orm:"column(date_created);type(datetime);auto_now_add"`
	DateModified time.Time `orm:"column(date_modified);type(datetime);auto_now"`
	Banned       int       `orm:"column(banned);"`
}
