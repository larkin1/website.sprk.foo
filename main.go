package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<!DOCTYPE html>
<html>
  <head><title>Sprk</title></head>
  <body><h1>Hi...</h1>
	but also what are you doing here? There is literally nothing lol</body>
</html>`)
	})

	http.ListenAndServe(":8080", nil)
}
