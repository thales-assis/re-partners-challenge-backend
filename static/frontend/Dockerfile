FROM nginx:1.27-alpine

COPY static/frontend/index.html /usr/share/nginx/html/index.html
COPY static/frontend/app.js /usr/share/nginx/html/app.js
COPY static/frontend/deploy/nginx/default.conf /etc/nginx/conf.d/default.conf
