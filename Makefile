mysql:
	docker run --name mysql-simple -e MYSQL_ROOT_PASSWORD=18032002 -p 3306:3306 -d mysql:8.0
dropmysql:
	docker stop mysql-simple && docker rm mysql-1
createdb:
	docker exec -it mysql-simple mysql -u root -p'18032002' -e "CREATE DATABASE simple_curd;"
dropdb:
	docker exec -it mysql-simple mysql-simple -u root -p'18032002' -e "DROP DATABASE simple_curd;"
migrateup:
	migrate -path db/migration -database "mysql://root:18032002@tcp(127.0.0.1:3306)/simple_curd?charset=utf8mb4&parseTime=True&loc=Local" -verbose up -force
migratedown:
	migrate -path db/migration -database "mysql://root:18032002@tcp(127.0.0.1:3306)/simple_curd?charset=utf8mb4&parseTime=True&loc=Local" -verbose down -all

.PHONY: mysql dropmysql  createdb dropdb migrateup migrateup
