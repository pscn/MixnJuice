FROM golang:1.11

# Prepare env
ENV GOPATH /go
ENV PATH ${GOPATH}/bin:$PATH
# Install dep (Note: run this locally to speed up image build)
#RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/github.com/pscn/MixnJuice/poc/go/mongo
WORKDIR /go/src/github.com/pscn/MixnJuice/poc/go/mongo
COPY . .
#RUN dep ensure
RUN go build
CMD ["./mongo"]
