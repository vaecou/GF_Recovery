
<!-- $env:CGO_ENABLED=0
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o main main.go -->


<!-- # 构建镜像并标记为 recovery
docker build -t recovery .

# 将 recovery 镜像重新打标签为 vaecou/recovery:latest
docker tag recovery vaecou/recovery:latest -->
