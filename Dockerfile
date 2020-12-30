FROM alpine:latest

RUN apk add --no-cache autoconf automake curl findutils fortify-headers gcc git libc-dev libtool make pkgconfig
RUN \
  git clone https://github.com/openvenues/libpostal && \
  cd libpostal && \
  ./bootstrap.sh && \
  ./configure --disable-data-download --prefix=/usr --datadir=/usr/share && \
  make && \
  make install && \
  libpostal_data download all /usr/share/libpostal
