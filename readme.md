# Problem #1
## To Run
- From the prebuild binary
```
$ ./mrcr-user server
```
- From source
```
$ cd merkari-user
$ go mod download
$ go run ./cli/user.go server
```
## Runtime Dependencies
- go 1.16
- mongoDB

# Problem #2
## To Run
- From the prebuild binary
```
$ ./mrcr-proxy server
```
- From source
```
$ cd merkari-proxy
$ go generate ./...
$ go mod download
$ go run ./cli/proxy.go server
```

## Build Dependencies
- protoc
- protoc-gen-gofast
## Runtime Dependencies
- go 1.16
- redis