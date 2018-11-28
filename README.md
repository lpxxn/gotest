# gotest
go test function

```
cd src
go test userinfo_test.go  userinfo.go utils.go -v
```

```
cd src/srctest/
go test userinfo_test.go
```

coverage
```
cd src
go test ./srctest/userinfo_test.go   -v -coverpkg=./
or
go test ./srctest/userinfo_test.go   -v -coverpkg ./

or
go test ./srctest/...  -v -coverpkg ./ -coverprofile cover.out
go tool cover -html=cover.out -o cover.html
```