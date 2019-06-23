package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://127.0.0.1:8080/:abc?dd=123")
	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", u.Scheme)
	fmt.Printf("%v\n", u.Host)
	fmt.Printf("%v\n", u.Path)
	fmt.Printf("%v\n", u.RawQuery)
	fmt.Printf("%v\n", u.RequestURI())
	fmt.Printf("%v\n", u.Port())
}
