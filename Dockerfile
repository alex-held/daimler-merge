# build stage
FROM golang AS builder
ARG VERSION=dev
ENV VERSION=$VERSION
LABEL maintainer = "contact@alexheld.io"

WORKDIR /go/src/github.com/alex-held/daimler-merge
COPY . ./
RUN export COMMIT=$(git rev-parse HEAD)
RUN go generate -tags=tools ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-extldflags -static -X=main.version=${VERSION} -X=main.commit=${COMMIT}" -a -o /app .


# final stage
FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]
