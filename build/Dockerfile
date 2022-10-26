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

RUN mkdir /entry
COPY ./entrypoint.sh /entry
RUN chmod a+x /entry

EXPOSE 80
ENTRYPOINT [ "/entry" ]


