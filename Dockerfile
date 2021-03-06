FROM golang:1.8
WORKDIR /usr/local/go/src/github.com/bgpat/dhcpd
ADD . ./
RUN go get ./... && go build --ldflags '-s -w -linkmode external -extldflags -static' -o /dhcpd

FROM scratch
COPY --from=0 /dhcpd /dhcpd
CMD ["/dhcpd"]
