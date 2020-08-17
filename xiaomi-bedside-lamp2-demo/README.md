```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build light-mapper.go
docker build -t lamp-mapper:v1.2.1 . 
docker tag lamp-mapper:v1.2.1 kindlaw/lamp-mapper:v1.2.1
docker push kindlaw/lamp-mapper:v1.2.1
```