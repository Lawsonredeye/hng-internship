# hng-internship

## Description
A simple golang webserver that returns a basic details of an hng intern (me).

## How to setup
To set up and run locally, there are some prerequisites to consider:
1. Go 1.23+ installed

### How to run
1. run `go mod tidy`
2. run `./main`

## DOCS
`curl -i localhost:8080/`

### Response
*Success 200 OK*

```JSON
{
  "email": "omoregbeeolawson@gmail.com",
  "created_time": "2025-01-29T20:07:12Z",
  "github_url": "https://github.com/lawsonredeye/hng-internship"
}
```

See more on golang:
https://hng.tech/hire/golang-developers