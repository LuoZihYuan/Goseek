
FROM golang:1.25.1-alpine AS goseek_dev
RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]


FROM goseek_dev AS build_stage
COPY . .
RUN swag init -g cmd/api/main.go -o docs
RUN CGO_ENABLED=0 GOOS=linux go build \
  -tags stage \
  -ldflags="-s -w" \
  -o /app/api \
  ./cmd/api

FROM goseek_dev AS build_prod
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
  -tags prod \
  -ldflags="-s -w" \
  -o /app/api \
  ./cmd/api


FROM alpine:latest AS goseek_stage
WORKDIR /app
COPY --from=build_stage /app/api .
COPY --from=build_stage /app/docs ./docs
COPY --from=build_stage /app/data ./data
RUN addgroup -g 1000 appuser && \
  adduser -D -u 1000 -G appuser appuser
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 8080
CMD ["./api"]


FROM alpine:latest AS goseek_prod
WORKDIR /app
COPY --from=build_prod /app/api .
COPY --from=build_prod /app/data ./data
RUN addgroup -g 1000 appuser && \
  adduser -D -u 1000 -G appuser appuser
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 8080
CMD ["./api"]