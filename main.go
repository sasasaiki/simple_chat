package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

func main() {
	r := newRoom()
	http.Handle("/", &templeteHandler{fileName: "chat.html"})
	//ルームを作る
	http.Handle("/room", r)
	//チャットルームを開始
	go r.run()

	//webServerの開始
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func (t *templeteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.fileName)))
		fmt.Println("テンプレートの読み込み")
	})

	e := t.templ.Execute(w, nil)

	if e != nil {
		fmt.Println("テンプレートの読み込みに失敗しています")
	}
}

type templeteHandler struct {
	once     sync.Once
	fileName string
	templ    *template.Template
}
