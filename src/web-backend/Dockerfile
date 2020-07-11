# 基础镜像
FROM node:12.14.0

# 支持中文
ENV LANF C.UTF-8

# 声明需要开放的端口
EXPOSE 8080

# 时区
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

RUN npm install -g serve

RUN mkdir /web-backend

COPY ./build /web-backend

WORKDIR /web-backend

CMD ["serve", "-s", "/web-backend", "-p", "8090"]