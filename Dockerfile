FROM scratch

ADD build/http /
ENTRYPOINT ["/http"]
