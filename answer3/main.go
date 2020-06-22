package main

import (
	"html/template"
	"net/http"
	"os"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	tmpl := `<doctype html>
<html>
<head>
<title>Hello {{.Name}}!</title>
</head>
<body>
Hello {{.Name}}!
</body>
</html>`

	t := template.Must(template.New("").Parse(tmpl))

	data := struct {
		Name string
	}{req.URL.Query().Get("name")}

	t.Execute(res, data)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":"+os.Args[1], nil)
}
