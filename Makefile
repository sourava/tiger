pwd:
	@echo $(PWD)

build:
	docker build -f Dockerfile -t tiger-dev .

up:
	@make build
	docker-compose -f docker-compose.yml up -d
	@make logs

e2e-tests:
	@mkdir -p e2e-results
	docker-compose -f docker-compose-e2e.yml down -v
	@make build
	docker-compose -f docker-compose-e2e.yml up -d
	docker run --network=tigerhall-e2e-test --mount type=bind,source=$(PWD)/e2e,target=/workdir/tests --mount type=bind,source=$(PWD)/e2e-results,target=/workdir/results ovhcom/venom:latest
	docker-compose -f docker-compose-e2e.yml logs -f

logs:
	docker-compose logs -f

down:
	docker-compose down -v

generate-swagger-doc:
	swag init -g app/server.go --output docs/