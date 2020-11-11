FROM golang:1.13 as builder
ENV GOPATH /go/
ENV PATH $GOPATH/bin:$PATH

ADD . $GOPATH/src/autoscaler
WORKDIR $GOPATH/src/autoscaler/cluster-autoscaler
RUN CGO_ENABLED=0 GOOS=linux go build --mod=vendor -o cluster-autoscaler .


FROM gcr.io/distroless/static:latest
COPY --from=builder /go/src/autoscaler/cluster-autoscaler/cluster-autoscaler /
CMD ["/cluster-autoscaler"]
