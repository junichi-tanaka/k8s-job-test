FROM alpine:latest

COPY testjob /testjob

ENTRYPOINT ["/testjob"]
