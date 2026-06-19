# syntax=docker/dockerfile:1

# 1) Build the React SPA.
FROM node:22-alpine AS web
WORKDIR /web
RUN corepack enable && corepack prepare pnpm@10.32.1 --activate
COPY web/package.json web/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
COPY web/ ./
RUN pnpm build

# 2) Build the server with the SPA embedded. No CGO → static binary.
FROM golang:1.26-alpine AS server
WORKDIR /src
ENV GOWORK=off CGO_ENABLED=0
COPY server/go.mod server/go.sum ./
RUN go mod download
COPY server/ ./
COPY --from=web /web/dist ./internal/webui/dist
RUN go build -trimpath -ldflags="-s -w" -o /out/server ./cmd/server

# 3) Minimal, non-root runtime image.
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=server /out/server /server
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/server"]
