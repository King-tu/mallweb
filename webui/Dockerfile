
# 基础镜像
FROM nginx
# 作者
MAINTAINER kingtu
# 删除原有配置文件
RUN rm /etc/nginx/conf.d/default.conf
# 复制新的配置文件
COPY ./conf.d/webui.conf /etc/nginx/conf.d
# 复制静态资源文件
COPY ./static /usr/share/nginx/html

