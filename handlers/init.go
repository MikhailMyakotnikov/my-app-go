package handlers

import (
	"database/sql"
	"html/template"
)

var db *sql.DB
var tpl *template.Template

func Init(database *sql.DB, templates *template.Template) {
	db = database
	tpl = templates
}
