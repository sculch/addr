FROM alpine:latest as libpostal
WORKDIR /libpostal
RUN apk add autoconf automake ca-certificates fortify-headers gcc git \
  libc-dev libtool make pkgconfig
RUN \
  git clone https://github.com/openvenues/libpostal . && \
  ./bootstrap.sh && \
  ./configure --disable-data-download --prefix=/usr --datadir=/usr/share && \
  make

FROM golang:alpine as builder
RUN apk add fortify-headers gcc git libc-dev make pkgconfig
COPY --from=libpostal /libpostal /libpostal
RUN cd /libpostal && make install
WORKDIR /app
COPY . /app
# RUN CGO_ENABLED=1 make test && CGO_ENABLED=1 make
RUN CGO_ENABLED=1 make

FROM alpine:latest
RUN apk add --no-cache ca-certificates curl fortify-headers findutils gcc \
  libc-dev make
COPY --from=libpostal /libpostal /libpostal
COPY --from=builder /app/addr /app/addr
RUN \
  cd /libpostal && \
  make install && \
  libpostal_data download all /usr/share/libpostal
WORKDIR /app

EXPOSE 8123
CMD ["/app/addr"]
