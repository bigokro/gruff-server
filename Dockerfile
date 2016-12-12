FROM scratch
ADD gruff-server /gruff-server
ENTRYPOINT ["/gruff-server"]
