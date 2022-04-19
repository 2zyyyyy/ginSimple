package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// sessions

// 初始化一个cookie存储对象
var store = sessions.NewCookieStore([]byte("test-secret"))

func main() {
	http.HandleFunc("/save", SaveSession)
	http.HandleFunc("/get", GetSession)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:", err)
		return
	}
}

func SaveSession(w http.ResponseWriter, r *http.Request) {
	// 获取一个session对象 session-name是session的名字
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 在session中存储值
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// 保存更改
	_ = session.Save(r, w)
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	foo := session.Values["foo"]
	fmt.Println(foo)
}
