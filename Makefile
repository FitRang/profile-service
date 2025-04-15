APP_NAME=profile-service
VERSION=v1.0.0
BUILD_TIME=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GIT_REVISION=$(shell git rev-parse --short HEAD)

build:
	go build -ldflags="-X main.buildRevision=$(GIT_REVISION) -X main.buildVersion=$(VERSION) -X main.buildTime=$(BUILD_TIME)" -o $(APP_NAME)

run: build
	./$(APP_NAME) --port :4444


db-init:
	export $$(cat .env | xargs) && \
	migrate -path db/migrations -database "$$POSTGRESQL_CONN_STRING" --verbose up
db-clean:
	export $$(cat .env | xargs) && \
	migrate -path db/migrations -database "$$POSTGRESQL_CONN_STRING" --verbose down
clean:
	rm -f $(APP_NAME)

version:
	@echo "Version: $(VERSION)"
	@echo "Git Revision: $(GIT_REVISION)"
	@echo "Build Time: $(BUILD_TIME)"

