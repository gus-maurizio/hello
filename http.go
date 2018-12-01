package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://bing.com/search/data?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n",u)
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Printf("%+v\n",u)
	fmt.Println(u)
}
