module github.com/panda8z/notebook

go 1.13

require github.com/gin-gonic/gin v1.5.0

replace (
	github.com/panda8z/notebook/db => ./dao/db
	github.com/panda8z/notebook/model => ./model
)
