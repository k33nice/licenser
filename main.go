package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
)

var templateDir = "templates"

const (
	Name    = "Alexandr Krykovliuk"
	Email   = "k33nice@gmail.com"
	Project = "Project"
	Port    = "33654"
)

// Copyright - handle name and email.
type Copyright struct {
	Name  string
	Email string
}

// Data - handle vars for licenes.
type Data struct {
	Year    int
	Project string
	Copyright
}

func main() {
	router := httprouter.New()
	router.GET("/:license", License)

	port, exsit := os.LookupEnv("PORT")
	if !exsit {
		port = Port
	}

	templatesBaseDir, exsit := os.LookupEnv("TEMPLATES")
	if exsit {
		templateDir = path.Join(templatesBaseDir, templateDir)
	}

	panic(http.ListenAndServe(net.JoinHostPort("localhost", port), router))
}

// License - return license.
func License(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	license := ps.ByName("license")
	re := regexp.MustCompile(`\-\d$`)
	if r := re.MatchString(license); r {
		license += ".0"
	}
	lt := path.Join(templateDir, license+".txt")

	if _, err := os.Stat(lt); os.IsNotExist(err) {
		Error(w)
		return
	}

	t, err := template.ParseFiles(lt)
	if err != nil {
		Error(w)
		return
	}

	n, e, y, p := Params(r.URL.Query())
	copy := Copyright{n, e}
	data := &Data{y, p, copy}

	if err := t.Execute(w, data); err != nil {
		Error(w)
		return
	}
}

// Error - return 404 http status.
func Error(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found.")
}

// Params - return query params if passed or defalts.
func Params(q url.Values) (string, string, int, string) {
	date := time.Now()
	Year, _, _ := date.Date()

	year, _ := strconv.Atoi(q.Get("y"))

	name := q.Get("n")

	email := q.Get("e")

	project := q.Get("p")

	return Default(name, Name).(string), Default(email, Email).(string),
		Default(year, Year).(int), Default(project, Project).(string)
}

// Default - retrun default for false.
func Default(v interface{}, def interface{}) interface{} {
	switch v.(type) {
	default:
		v = def
	case bool:
		if v == false {
			v = def
		}
	case int:
		if v.(int) <= 0 {
			v = def
		}
	case string:
		if v == "" {
			v = def
		}
	}

	return v
}
