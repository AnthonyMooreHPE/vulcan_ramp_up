FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./vulcan_ramp_up ./hello.go


FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/vulcan_ramp_up .
EXPOSE 7000
ENTRYPOINT ["./vulcan_ramp_up"]