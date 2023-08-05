build:
	docker build -f Dockerfile -t tiger-dev .

up:
	@make build
	docker-compose -f docker-compose.yml up -d
	@make logs

logs:
	docker-compose logs -f

down:
	docker-compose down -v
