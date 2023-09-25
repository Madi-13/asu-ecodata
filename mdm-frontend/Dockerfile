FROM reg.appleseed.kz/docker-proxy/library/nginx:stable
ARG AUTH_URL
ARG REFERENCES_URL

RUN echo "👋 Hi, i am a docker"

COPY ./dist /usr/share/nginx/html
COPY ./nginx.conf /temporary-nginx.conf

RUN envsubst '${AUTH_URL},${REFERENCES_URL}' < /temporary-nginx.conf > /etc/nginx/conf.d/default.conf
RUN cat /etc/nginx/conf.d/default.conf

EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
