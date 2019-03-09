package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
				    <title>チャットしようよ</title>
				</head>
				<body>
				    チャットだ！！
				</body>
			</html>
		`))

		// Start webserver
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("ListenAndServer:", err)
		}
	})
}
