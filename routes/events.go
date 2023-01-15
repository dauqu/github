package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Events(w http.ResponseWriter, r *http.Request) {
	//Print all body request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	//Print all header request
	header := r.Header

	//Print all query request
	query := r.URL.Query()

	//Print all cookie request
	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
	}

	//Print all form request
	form := r.FormValue("token")

	//Print all method request
	method := r.Method

	fmt.Println("Body:", string(body))
	fmt.Println("Header:", header)
	fmt.Println("Query:", query)
	fmt.Println("Cookie:", cookie)
	fmt.Println("Form:", form)
	fmt.Println("Method:", method)
}

