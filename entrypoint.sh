#!/bin/bash
set -e

certbot certonly --nginx --non-interactive --agree-tos -m your-email@dev.e-frontend.ru -d dev.e-frontend.ru

nginx -g 'daemon off;'
