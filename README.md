# SHRT Go

Simple URL shortener api written in Golang

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
