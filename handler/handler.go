package handler

import (
	"kumparan/config"
	"kumparan/config/database"
	article "kumparan/module/v1/user"
	authMid "kumparan/utl/middleware/auth"
)

type Service struct {
	MiddlewareAuth *authMid.Handle
	ArticleModule  *article.Module
}

func InitHandler() *Service {

	// mysql init
	MySQLConnection, err := database.MysqlDB()
	if err != nil {
		panic(err)
	}

	config := config.Configuration{
		MysqlDB: MySQLConnection,
	}

	// set service modular
	middlewareAuth := authMid.InitAuthMiddleware(config)
	moduleArticle := article.InitModule(config)

	return &Service{
		ArticleModule:  moduleArticle,
		MiddlewareAuth: middlewareAuth,
	}
}
