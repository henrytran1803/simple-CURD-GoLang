# simple-CURD-GoLang
## Framework
- GIN
- GORM
- Migrate
## MakeFile
$brew install golang-migrate
$migrate create -ext sql -dir db/migration -seq init_schema
- create docker mysql
- create db
- dropdb
- create migrate
- drop migrate
## Clean architecture
![](img_1.png)
### SQL - MYSQL
Gồm 3 bảng cơ bản roles, users, và account
roles many to one với users, và users one to one account
**Diagram**
![](img.png)
### Deloy Docker