# CXRUS Golang API

This is Cxrus Golang API created for learning purposes

To run this project you must install the golang from this website https://golang.org/dl/

After downloading and installing you must set your **GOPATH**

You can find instruction of setting the **GOPATH** in here https://github.com/golang/go/wiki/SettingGOPATH#bash

When the **GOPATH** already defined in your system, please create folder **src** on the **GOPATH** and clone this project on **src** folder

If the project has an error in import some of package it means the package isn't in **%GOPATH%**/**src** you must install the package by running this command

`go get ${package_name}`

## **Example**

`go get github.com/dgrijalva`

It will download the package put in **%GOPATH%/src** 

If there's no error when compiling the code you can run the project by going to directory of your application example like **%GOPATH%/src/cxrus-app/** and run the main program of your golang project by using this command

`go run ${main_program}`

Basically main program that can be executable by go run syntax is the .go extension that has **func main()** inside the file

## Example
~~~
func main() {

db.InitDB()
log.Println("Initializing DB Success")

router := routes.InitBaseRoutes()
log.Println("All Routes Has Been Initialized")
err := http.ListenAndServe(":"+"8080", router)
if err != nil {
    log.Fatal(err)
} else {
    log.Println("Server is start using port 3000")
}

}
~~~