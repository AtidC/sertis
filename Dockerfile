#=============================================================
#--------------------- build stage ---------------------------
#=============================================================
FROM golang:1.17.6-alpine3.15 AS build_stage
LABEL stage=build_stage

ENV PACKAGE_PATH=blog

RUN mkdir -p /go/src/
WORKDIR /go/src/$PACKAGE_PATH

ENV PATH="/usr/local/go/bin:$PATH"

COPY . /go/src/$PACKAGE_PATH/
RUN go mod download

RUN go build -o bin/blog

#=============================================================
#--------------------- runtime stage -------------------------
#=============================================================
FROM alpine:3.15 AS final_stage
LABEL stage=final_stage

ENV PACKAGE_PATH=blog

COPY --from=build_stage /go/src/$PACKAGE_PATH/bin/blog /go/src/$PACKAGE_PATH/
COPY --from=build_stage /go/src/$PACKAGE_PATH/config /go/src/$PACKAGE_PATH/config
RUN mkdir /go/src/$PACKAGE_PATH/tmp 

WORKDIR /go/src/$PACKAGE_PATH/

ENTRYPOINT ./blog dev
EXPOSE 80
