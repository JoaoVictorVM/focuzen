WEBUI_DIST := server/internal/webui/dist
IMAGE ?= focuzen:local

.PHONY: web embed build test lint clean docker-build docker-run

# Build the SPA and copy it into the server's embed directory.
web:
	cd web && pnpm install && pnpm build
	rm -rf $(WEBUI_DIST)
	mkdir -p $(WEBUI_DIST)
	cp -r web/dist/. $(WEBUI_DIST)/

# Alias for "the embed dir is ready".
embed: web

# Full build: SPA embedded into the server binary.
build: web
	cd server && go build -o ../bin/server ./cmd/server

test:
	cd server && go test ./...
	cd web && pnpm test

lint:
	cd server && golangci-lint run ./...
	cd web && pnpm typecheck

clean:
	rm -rf web/dist $(WEBUI_DIST) bin

# Build the production container image (SPA + server, distroless, non-root).
docker-build:
	docker build -t $(IMAGE) .

# Run the image locally (expects YOUTUBE_API_KEY in the environment).
docker-run:
	docker run --rm -p 8080:8080 -e YOUTUBE_API_KEY=$(YOUTUBE_API_KEY) $(IMAGE)
