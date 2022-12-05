.PHONY: docker-compose createdb

docker-compose:
	sudo docker-compose up -d  --build

createdb:
	sudo docker exec -it erajaya-testcoding_postgres_1 createdb --username=user --owner=user db_product


