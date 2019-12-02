FROM scratch

COPY gopath/bin/testjob /testjob

ENTRYPOINT ["/testjob"]
