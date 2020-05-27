package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (
	coockieDelimiter = "|"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	sessionCookie := getUserSessionCookie(w, req)

	//counterCookie := getUserVisitCouterCookie(w, req)

	//newCoockieValues := appendValueCookie(sessionCookie, w)

	//fmt.Printf("Sesssion ID cookie is the following: %v and CounterCookie is the following: %v", sessionCookie.Value, counterCookie.Value)

	if req.Method == http.MethodPost {
		multiPartFile, fileHeader, err := req.FormFile("nf")

		if err != nil {
			log.Print(err)
		}
		defer multiPartFile.Close()

		// create sha for file name
		extension := strings.Split(fileHeader.Filename, ".")[1]
		hash := sha1.New()
		io.Copy(hash, multiPartFile)
		fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + extension

		// create new file

		wd, err := os.Getwd()

		if err != nil {
			log.Println(err)
		}

		path := filepath.Join(wd, "public", "pics", fname)

		newFile, err := os.Create(path)

		if err != nil {
			log.Println(err)
		}

		defer newFile.Close()

		//copy
		multiPartFile.Seek(0, 0)
		io.Copy(newFile, multiPartFile)

		// add filename to user's cookie

		sessionCookie = appendPredefinedValueToCooki(sessionCookie, w, fname)

	}

	xs := strings.Split(sessionCookie.Value, coockieDelimiter)

	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getUserVisitCouterCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	ck, err := req.Cookie("visitCounter")

	if err == http.ErrNoCookie {

		ck = &http.Cookie{
			Name:  "visitCounter",
			Value: "0",
		}
	}
	cookieValue, _ := strconv.Atoi(ck.Value)

	cookieValue++

	ck.Value = strconv.Itoa(cookieValue)

	http.SetCookie(w, ck)

	return ck
}

func getUserSessionCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {

	ck, err := req.Cookie("sessionId")

	if err == http.ErrNoCookie {
		sID, _ := uuid.NewV4()
		coockieExpireTime := 10

		ck = &http.Cookie{
			Name:   "sessionId",
			Value:  sID.String(),
			MaxAge: coockieExpireTime,
		}

		ck.Value += "|"
		http.SetCookie(w, ck)
	}

	return ck
}

func appendPredefinedValueToCooki(ck *http.Cookie, w http.ResponseWriter, name string) *http.Cookie {
	s := ck.Value

	if !strings.Contains(s, name) {
		s += "|" + name
	}

	ck.Value = s

	http.SetCookie(w, ck)

	return ck
}

func appendValueCookie(ck *http.Cookie, w http.ResponseWriter) []string {

	picNames := strings.Split(ck.Value, coockieDelimiter)

	if strings.Contains(ck.Value, coockieDelimiter) {

		if len(picNames) < 3 {
			i := 0
			for len(picNames) < 5 {
				fmt.Println(len(picNames))
				picName := "picNumber" + strconv.Itoa(len(picNames))
				ck.Value += picName + "|"
				picNames = append(picNames, picName)
				i++
			}

			http.SetCookie(w, ck)
		}
	}

	return picNames
}
