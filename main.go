package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UserInfo struct {
	Email string `json:"email"`
	CurrentDatetime string `json:"current_datetime"`
	GithubUrl string `json:"github_url"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		userUrl := "https://github.com/lawsonredeye/hng-internship"
		t := time.Now().UTC()
		ft := t.Format(time.RFC3339)
		user := UserInfo{"omoregbeeolawson@gmail.com",
			ft,
			userUrl,
		}
		w.Header().Add("Content-Type", "application/json")
		b, _ := json.Marshal(user)
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}