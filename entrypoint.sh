#!/bin/bash
set -e

certbot certonly --nginx --non-interactive --agree-tos -m d_maximyuk@icloud.com -d dev.e-frontend.ru

nginx -g 'daemon off;'
