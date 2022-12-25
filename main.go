package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Name struct{
	FirstName string
	LastName string
}

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		method:=r.Method
		if method=="GET"{
			tmpl,_:=template.ParseFiles("form.html")
			tmpl.Execute(w,nil)
		}
		 if method=="POST"{

			firstname:=r.FormValue("first")
			lastname:=r.FormValue("last")
			fmt.Println(firstname+" "+lastname)

			tmpl,_:=template.ParseFiles("result.html")

			name:=Name{
				FirstName: firstname,
				LastName: lastname,
			}
			err:=tmpl.Execute(w,name)
			if err!=nil{
				log.Fatal(err)
			}
		}
	})
	http.ListenAndServe(":8000",nil)

}