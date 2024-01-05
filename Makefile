init:
	go mod tidy
	go mod vendor

run:
	go run main.go

build:
	go build -o bin/dating main.go

test:
	@go test ./...

test-coverage:
	@go test ./... -coverprofile=cover.out.tmp
	@cat cover.out.tmp | grep -v "_mock.go" > cover.out
	@go tool cover -html=cover.out -o cover.html
	@open cover.html

build-docker:
	docker-compose up --build -d

generate-mock:
	@echo "GENERATING ..."
	@echo "- jwt"
	@mockgen -destination=pkg/auth/jwt/jwt_mock.go -package=jwtAuth -source=pkg/auth/jwt/jwt.go
	@echo "- middleware"
	@mockgen -destination=pkg/auth/middleware/middleware_mock.go -package=middleware -source=pkg/auth/middleware/middleware.go
	@echo "- pagination"
	@mockgen -destination=helpers/pagination/pagination_mock.go -package=pagination -source=helpers/pagination/pagination.go
	@echo "[API]"
	@echo "1. AUTH"
	@mockgen -destination=api/domains/interfaces/auth_mock.go -package=interfaces -source=api/domains/interfaces/auth.go
	@echo "DONE ..."
