package main

import (
	"fmt"
	"net/url"
)

func main() {
	base_url := "https://api.weather.com/data?city=delhi&unit=metric"
	u, _ := url.Parse(base_url)
	fmt.Println("u: ", u)
	params := u.Query()
	params.Set("unit", "imperial") //modify
	params.Set("lang", "en")       // add
	u.RawQuery = params.Encode()   //Without u.RawQuery = params.Encode(), the final URL won’t be updated with your changes.
	fmt.Println("updated url: ", u.String())
}
