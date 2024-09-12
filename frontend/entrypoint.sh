#!/bin/sh

# Заменяем плейсхолдеры в config-template.js на реальные значения переменных окружения
envsubst < /usr/share/nginx/html/config-template.js > /usr/share/nginx/html/config.js

# Запускаем Nginx
nginx -g 'daemon off;'