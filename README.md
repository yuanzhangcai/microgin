# 微服务gin后台
### 微服务gin后台后台框架以go-micro + go-gin组成，程序支持两种运行方式：
#### 1、微服务型式启动
#### 2、go-gin web服务器型式启动
### 配置文件使用了micro/config模块，支持配置文件动态更新，配置文件修改之后，不需要重新启动程序；支持多配置文件加载，后加载的配置文件中的配置参数，如果与之前加载的配置文件参数相同，则会覆盖掉之前配置。
### 日志使用了sirupsen/logrus模块。

## 编译

### 建议通make命令编译

```
$ make prod     # 编译正式环境程序

$ make test     # 编译测试环境程序，代码执行时会多加载一次测试环境配置文件(./etc/test.toml)

$ make dev      # 编译本地开发环境程序，代码执行时会多加霜一次本地环境配置文件(./etc/dev.toml)

$ make install  # 安装gowatch，执行该命令会给go.mod添加gowatch相关的依赖包，建议到工作目录外手动执行 go get github.com/silenceper/gowatch

$ make watch    # 用gowatch启动程序，默认的运行环境为local。以gowatch启动程序，可以实现热编译，当文件有修改时，
                # 程gowatch会自动编译运行程序，不需要每次手动编译运行。适合本地开发测试。正式环境勿用。
```

## 运行

### 微服务后台，支持两种运行方式，具体以哪种方式运行，由配置文件./etc/nnc.toml中的modes参数决定

### 1、微服务型式启动

```
$ micro api --handler=http

$ ./microgin 
```

### 2、gin web服务启动
```
$ ./microgin
```

## 测试
### 运行如下代码：
```
$ curl 'http://localhost:8080/microgin/version'
```

### 返回结果：
```
{
    "data": {
        "buildby": "zacyuan",
        "buildtime": "2019-12-31T10:29:39",
        "commit": "56e61f3a41e6c9e8681d6d7fe8f5b990db57147f",
        "env": "dev",
        "version": "V0.0.1"
    },
    "msg": "OK",
    "ret": 0
}
```