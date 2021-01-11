package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func middleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method != "HEAD" && r.Method != "GET") && r.UserAgent() != "lic engine" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		encryptedKey := r.Header.Get("X-USER-KEY")
		if encryptedKey != "" {
			key := encryptDecrypt(encryptedKey, "X-USER-KEY")
			s := findSub(key)
			if !s.isZero() {
				r.Header.Set("X-USER-KEY", key)
				f(w, r)
				return
			}
		} else {
			fmt.Fprintf(w, "%s", time.Now().Format(time.RFC3339))
		}
	}
}

func libHost(w http.ResponseWriter, r *http.Request) {
	libContent := getLibByName(r.Header.Get("X-USER-LIB"))
	if libContent != nil {
		w.Write(libContent)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func check(w http.ResponseWriter, r *http.Request) {
	s := findSub(r.Header.Get("X-USER-KEY"))
	log.Println(s.Name, "logged in with ip:", r.RemoteAddr)
	content, _ := json.Marshal(s)
	w.Header().Set("X-USER-RESPONSE", string(content))
}

func softwareList(w http.ResponseWriter, r *http.Request) {
	s := findSub(r.Header.Get("X-USER-KEY"))
	json.NewEncoder(w).Encode(getAvailablesWareList(s))
}
