::windows环境下编译linux环境下运行程序
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go