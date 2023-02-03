#请选择golang版本
BUILD_IMAGE_SERVER  = golang:1.19.5

# 打包可执行文件lms_server
build:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go -o lms_server