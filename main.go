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
	"strings"
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
	Name, Email string
}

// Years - handle copyright years.
type Years struct {
	From, To int
}

// Data - handle vars for licenes.
type Data struct {
	Project string
	Years
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
	data := &Data{p, y, copy}

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
func Params(q url.Values) (string, string, Years, string) {
	date := time.Now()
	Year, _, _ := date.Date()

	years := q.Get("y")
	yearsList := strings.SplitN(years, "-", 2)

	name := q.Get("n")

	email := q.Get("e")

	project := q.Get("p")

	n, ok := Default(name, Name).(string)
	if !ok {
		n = name
	}
	e, ok := Default(email, Email).(string)
	if !ok {
		e = email
	}
	y, ok := Default(yearsList, Year).(Years)
	if !ok {
		y = Years{To: Year}
	}
	p, ok := Default(project, Project).(string)
	if !ok {
		p = project
	}

	return n, e, y, p
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
	case []string:
		years, ok := v.([]string)
		if !ok || len(years) == 0 || years[0] == "" {
			v = Years{To: def.(int)}
			break
		}
		y := Years{}
		y.To, _ = strconv.Atoi(strings.Trim(years[0], " "))
		if len(years) > 1 {
			y.From, _ = strconv.Atoi(strings.Trim(years[0], " "))
			y.To, _ = strconv.Atoi(strings.Trim(years[1], " "))
		}
		v = y
	}

	return v
}
