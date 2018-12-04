# gotest

specify test file
```
cd app1/api
go test -v user_test.go  user.go
```

```
cd app1
go test -v utils/tool_test.go utils/tool.go
```

test all file in package
```
cd app1
go test -v
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