FROM node:16.13.2 as builder

WORKDIR /app
COPY ./ui /app

RUN npm ci

RUN npm run build


FROM centos:7 as hwameistor

COPY ./_build/hwameistor-ui /hwameistor-ui

FROM nginx:1.21.6-alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY --from=builder /app/nginx.conf /etc/nginx/nginx.conf
COPY --from=hwameistor /hwameistor-ui /hwameistor-ui

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

COPY /server/config/application.yaml /application.yaml
RUN chmod +x /application.yaml

EXPOSE 80
ENTRYPOINT [ "/entrypoint.sh" ]


