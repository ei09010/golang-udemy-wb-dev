package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", DealCookie)
	http.HandleFunc("/delete", ClearCookie)
	http.ListenAndServe(":8080", nil)
}

func DealCookie(w http.ResponseWriter, req *http.Request) {

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

		fmt.Fprintf(w, "cookie counter updated to %v", cookieValue)

}


func ClearCookie(w http.ResponseWriter, req *http.Request){
	
	ck, err := req.Cookie("visitCounter")

	if err == http.ErrNoCookie {

		ck = &http.Cookie{
			Name:  "visitCounter",
			Value: "1",
		}

		fmt.Fprintf(w, "created my cookie for the first time")
	}

	ck.MaxAge = -1

	http.SetCookie(w, ck)

	fmt.Fprintf(w, "cookie deleted")

}