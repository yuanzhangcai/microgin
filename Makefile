default: dev

PACKAGE="github.com/yuanzhangcai/microgin/common"

USER=`whoami`
GIT_TAG=`git describe --tags`
GIT_COMMIT=`git rev-parse HEAD`
BUILD_TIME=`date '+%F %T'`
VERSION=`git rev-list --tags --max-count=1 | xargs git describe --tags`
GO_VERSION=`go version`
LDFLAGS="-X ${PACKAGE}.Commit=${GIT_COMMIT} -X '${PACKAGE}.BuildTime=${BUILD_TIME}' -X ${PACKAGE}.BuildUser=${USER} -X ${PACKAGE}.Version=${VERSION} -X '${PACKAGE}.GoVersion=${GO_VERSION}' "

# 安装gowacth，执行该命令会给go.mod添加gowatch相关的依赖包，建议到工作目录外手动执行 go get
install:
	go get github.com/silenceper/gowatch

# 热编译启动程序 默认指定运行环境为本地运行环境，配置详见gowatch.yml
watch:
	gowatch

# 编译本地开发运行环境程序
dev:
	go build -ldflags ${LDFLAGS}" -X ${PACKAGE}.Env=dev" -race

# 编译测试运行环境程序
test:
	go build -ldflags ${LDFLAGS}" -X ${PACKAGE}.Env=test" -race

# 编译正式运行环境程序
prod:
	go build -ldflags ${LDFLAGS}" -X ${PACKAGE}.Env=prod"

.PHONY: install watch dev test prod