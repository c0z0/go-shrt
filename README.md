[![Deploy to now](https://deploy.now.sh/static/button.svg)](https://deploy.now.sh/?repo=https://github.com/zpnk/hello-world)
# [Go SHRT](https://gs.cserdean.me)

Simple URL shortener api written in Golang

## Getting started
```
$ git clone https://github.com/c0z0/shrt-go && cd shrt-go
$ MONGO_URI="mongodb://<user>@<password>0000.mlab.com/<db-name>" MONGO_DB="<dbname>" go run main.go
```

## Routes

### Shorten
`POST /s`

```json
{
	"url": "https://example.com"
}
```

Response

```json
{
	"url": "https://example.com",
	"id": "JiBBe"
}
```

### Go to url

`GET /{id}`

Redirects to shortened url
