package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"
	"text/template"

	// _ "github.com/mattn/go-sqlite3"
)

func RootHandler(res http.ResponseWriter, req *http.Request, report Report) {
	tmpl, _ := ioutil.ReadFile("views/index.html")

	t := template.New("").Funcs(funcMap)
	t, err := t.Parse(string(tmpl))
	checkError(err)

	err = t.Execute(res, report.Run())
	checkError(err)
}

func StaticHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, req.URL.Path[1:])
}

/*********************** TEMPLATE FUNC MAP ***********************/

var funcMap = template.FuncMap{
	"rate":    rate,
	"round":   round,
	"comma":   comma,
	"acronym": acronym,
	"crop":    crop,
	"month":   month,
}

func rate(n, m int) float64 {
	return round((float64(n)/float64(m))*100, 1)
}

func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor((f*shift)+0.5) / shift
}

func comma(list []string) (results string) {
	results = list[0]
	for _, v := range list[1:] {
		results += ", " + v
	}
	return
}

func acronym(title, delim string) (results string) {
	segments := strings.Split(title, " ")
	for _, v := range segments {
		results += fmt.Sprintf("%s%s", strings.ToUpper(v[:1]), delim)
	}
	return
}

func crop(title string, spaces int) (results string) {
	return title[spaces:]
}

func month(m int) string {
	return map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}[m]
}

/*********************** ERROR CHECKING ***********************/

func checkError(err error) {
	if err != nil {
		fmt.Println("\nERR: ", err.Error())
		os.Exit(1)
	}
}
