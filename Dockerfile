# # 编译镜像
# FROM golang as builder

# ENV GOPROXY="https://goproxy.io" GOSUMDB="off" GOPRIVATE="github.com/yuanzhangcai"

# # 指定工作目录
# WORKDIR /data/microgin/

# # 拷贝文件
# ADD . .

# # 下载依赖
# RUN go mod download && make dev

# # 运行程序镜像
# FROM scratch
# # FROM ubuntu

# WORKDIR /data/microgin/

# RUN mkdir -p /data/logs && chmod 777 /data/logs

# # 拷贝文件
# COPY --from=builder /data/microgin/microgin .
# COPY --from=builder /data/microgin/config ./config

# # 开放端口
# EXPOSE 8080

# # 执行程序
# CMD ["/data/microgin/microgin"]

# ########################################## 上面纯docker镜像编译，每次编译都要重新下载一次go mod，编译太慢。 #######################################################

# 运行程序镜像
# FROM scratch  #空镜像，涉及到创建log目录，log目录加权限，所以不适用。
FROM alpine
# FROM ubuntu

WORKDIR /data/microgin/

RUN mkdir -p /data/logs && chmod 777 /data/logs

# 拷贝文件
COPY ./microgin .
COPY ./config ./config

# 开放端口
EXPOSE 8080

# 执行程序
CMD ["/data/microgin/microgin"]

#编译docker镜像
#docker build -t microgin .