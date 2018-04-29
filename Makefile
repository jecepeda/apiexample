test:
	docker-compose run web buffalo test
build:
	docker-compose build
create:
	docker-compose run web buffalo db create
up:
	docker-compose up
