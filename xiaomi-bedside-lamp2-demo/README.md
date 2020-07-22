```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build light-mapper.go
docker build -t lamp-mapper:v1.2 . 
```