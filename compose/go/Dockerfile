FROM golang:1.19.1-alpine

WORKDIR /app
COPY . .

RUN apk add build-base &&\
    make build


# RUN addgroup --system todd &&\
#     adduser --system --ingroup todd todd &&\
#     chown todd:todd .

# USER todd

CMD ["main"]