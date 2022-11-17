set GOPATH=D:\gopath
set GOARCH=amd64
set GOOS=linux

@REM cd zinx-demo
@REM go build main.go
@REM move main bin/zinx-demo

@REM cd redis-find-bigkey
@REM go build -o redis-find-nottl  main.go


cd redis-key-parser
go build -o redis-token-nottl-delete  main.go

pause
