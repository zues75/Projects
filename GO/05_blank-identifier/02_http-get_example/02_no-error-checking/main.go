package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.mcleods.com")
	page, _ := iountil.ReadAll(res.Body)
	res.Body.Close()
	page.Body.Close()
	fmt.Println("%s", page)
}
