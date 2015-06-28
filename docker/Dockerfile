FROM gliderlabs/alpine

MAINTAINER Antons Kranga <akranga@gmail.com>

RUN apk-install nginx
COPY nginx.conf /etc/nginx/nginx.conf	
COPY tmp/consul-template /usr/sbin/consul-template
# run build.sh file to download consul-template
RUN chmod +x /usr/sbin/consul-template

ENV SERVICE_NAME proxy

EXPOSE 80

CMD ["/usr/sbin/nginx"]