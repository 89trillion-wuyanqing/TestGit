package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var HandleHelloWorld = func(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body>Hello World!</body></html>")
}

func main() {

	/*user :=model.Users{}
	user.SetId("1234")
	user.Age=22
	user.Name="吴延卿"
	user.Phone="12345678901"
	UserMap , err :=service.InsertUsers(user)

	user2 :=model.Users{}
	user2.SetId("12345")
	user2.Age=23
	user2.Name="吴延卿2"
	user2.Phone="123456789012"
	UserMap,err =service.InsertUsers(user2)
	fmt.Println(UserMap)
	if err != nil {
		fmt.Println(err)
	}

	UserMap,err = service.DeleteUsers("1234")
	fmt.Println(UserMap)*/

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	HandleHelloWorld(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("34556")
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	fmt.Println("1234")
}
