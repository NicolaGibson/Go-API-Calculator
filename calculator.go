package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

//http://localhost:3000/add?a=1&b=2
func main(){
	http.HandleFunc ("/add", Add)
	http.HandleFunc ("/minus", Minus)
	http.HandleFunc ("/multiply", Multi)
	http.HandleFunc ("/divide", Divide)
	http.ListenAndServe(":3000", nil)
}


func Add (w http.ResponseWriter, r *http.Request){

	str1 := r.URL.Query().Get("a")
	str2:= r.URL.Query().Get("b")

	int1, err := strconv.ParseFloat(str1, 32); err = nil
	fmt.Println(reflect.TypeOf(int1))
	if err!= nil{
	log.Fatal(err, "Please select numbers between x and x")
}

	int2, err := strconv.ParseFloat(str2, 32); err = nil
	fmt.Println(reflect.TypeOf(int2))
	if err!= nil{
		log.Fatal(err)
	}

		fmt.Fprint(w,add(float32(int16(int1)), float32(int16(int2))))

}

func Minus (w http.ResponseWriter, r *http.Request){

	str1 := r.URL.Query().Get("a")
	str2:= r.URL.Query().Get("b")

	int1, err := strconv.Atoi(str1)
	if err!= nil{
		log.Fatal(err)
	}

	int2, err:= strconv.Atoi(str2)
	if err!= nil{
		log.Fatal(err)
	}

	fmt.Fprint(w,minus(int1,int2))
}

func Multi (w http.ResponseWriter, r *http.Request){

	str1 := r.URL.Query().Get("a")
	str2:= r.URL.Query().Get("b")

	int1, err := strconv.Atoi(str1)
	if err!= nil{
		log.Fatal(err)
	}

	int2, err:= strconv.Atoi(str2)
	if err!= nil{
		log.Fatal(err)
	}

	fmt.Fprint(w,multi(int1, int2))

}
func Divide (w http.ResponseWriter, r *http.Request) {
	str1 := r.URL.Query().Get("a")
	str2 := r.URL.Query().Get("b")

	int1, err := strconv.Atoi(str1)
	if err != nil {
		log.Fatal(err)
	}
	int2, err := strconv.Atoi(str2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w,divide(int1, int2))
}



func add (a float32,  b float32) float32{
	//  if cannot be added together - return sum or return error
	sum := a + b
		return sum

}

func minus (a int, b int) int{
	sum := a - b
	return sum
}

func multi (a int, b int) int{
	sum := a * b
	return sum
}

func divide (a int, b int) int{
	sum := a /b
	return sum
}

/*
pass in 2 as a first example and then 2 as second example can you make 4.
http:localhost:8080/2/2
4
http:localhost:8080/4/2
6
*/

/*handle errors and floats, google url parsing, write tests to make sure operations
performing correctly.  Figure out how to take data from the browser and use that in code. More operations e.g.
to the power of.*/