module github.com/panda8z/notebook

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmoiron/sqlx v1.2.0
)

replace (
	github.com/panda8z/notebook/db => ./dao/db
	github.com/panda8z/notebook/model => ./model
)
