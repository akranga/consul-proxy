#!/bin/bash
mkdir tmp
rm -rf tmp/*
curl -o tmp/consul-template.tar.gz -L -J https://github.com/hashicorp/consul-template/releases/download/v0.10.0/consul-template_0.10.0_linux_amd64.tar.gz
tar -xvzf tmp/consul-template.tar.gz -C tmp
mv tmp/consul-template_*/consul-template tmp
rm -rf mv tmp/consul-template_*
docker build -t akranga/nginx .
