# gotest

specify test function
```
go test -v -run=TestGetOrderList
go test -v -run="TestGetOrderList|TestNewUserInfo"
```

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
cd api
go test -v
```

test all subpackages recursively
```
cd app1
go test -v ./...
```

coverage test specify file
```
cd app1
go test -v -coverprofile cover.out ./api/user_test.go ./api/user.go
go tool cover -html=cover.out -o cover.html
```

coverage test all file in package
```
cd app1
go test -v -coverprofile cover.out ./api/...
go tool cover -html=cover.out -o cover.html
```