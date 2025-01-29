package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UserInfo struct {
	Email string `json:"email"`
	CurrentDatetime time.Time `json:"created_time"`
	GithubUrl string `json:"github_url"`
}

func main() {
	// returns the details in the format:
	// {
	//   "email": "your-email@example.com",
	//   "current_datetime": "2025-01-30T09:30:00Z",
	//   "github_url": "<https://github.com/yourusername/your-repo>"
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		userUrl := "https://github.com/lawsonredeye/hng-internship"
		user := UserInfo{"omoregbeeolawson@gmail.com",
			time.Now().UTC(),
			userUrl,
		}
		w.Header().Add("Content-Type", "application/json")
		b, _ := json.Marshal(user)
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}