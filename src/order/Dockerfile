# 基础镜像
FROM debian:9

# 支持中文
ENV LANF C.UTF-8

# 声明需要开放的端口
EXPOSE 50052

# 时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

RUN mkdir /order

COPY ./config.yaml /order

COPY ./order /order

WORKDIR /order

ENTRYPOINT ["./order"]