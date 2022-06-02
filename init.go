package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// path file is depends to enveronment.
func templ() *Template {
	var p string
	//if os.Getenv("USERNAME") != "fedor" {
	//	p = "/root/store/"
	//}
	files := []string{
		p + "tmpl/home.html", p + "tmpl/sign.html", p + "tmpl/login.html",
		p + "tmpl/updatefotos.html", p + "tmpl/404.html", p + "tmpl/acount.html",
		p + "tmpl/upload.html", p + "tmpl/upacount.html", p + "tmpl/messages.html",
		p + "tmpl/part/header.html", p + "tmpl/part/footer.html", p + "tmpl/activity.html",
	}
	return &Template{templates: template.Must(template.ParseFiles(files...))}
}

// folder when photos is stored.

func photoFold() string {
	//if os.Getenv("USERNAME") == "fedor" {
	//	return "/home/fedor/repo/files/"
	//}
	return "../files/"
}

// where assets  path ?
func assets() string {
	home, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	if home != "/Users/fedora/repo/social" {
		return "/root/social/assets"
	}
	return "assets"
}

var db *sql.DB

func setdb() *sql.DB {
	db, err := sql.Open(
		"mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {

		log.Println("open database error: ", err)
		switch {
		case strings.Contains(err.Error(), "connection refused"):
			// TODO handle errors by code of error not by strings.

			//cmd := exec.Command("mysql.server", "restart")
			// for systemd linux : exec.Command("sudo", "service", "mariadb", "start")
			//cmd.Stdin = strings.NewReader(os.Getenv("JAWAD"))
			//err = cmd.Run()
			if err != nil {
				fmt.Println("error when run database cmd ", err)
			}
		default:
			log.Println("not knowen err at db.Ping() func")
			log.Println("unknown this error", err)
			os.Exit(1)
			//return nil
		}
	}
	return db
}
