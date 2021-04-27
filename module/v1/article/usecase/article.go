package usecase

import (
	"kumparan/config"
	"kumparan/model"
	"kumparan/module/v1/article/repo"
)

// get list article
func ArticleList(conf config.Configuration) (users []model.Article, err error) {
	db := conf.MysqlDB
	return repo.GetUserList(db)
}

// get detail article
func ArticleDetail(conf config.Configuration, userId int) (users model.Article, err error) {
	db := conf.MysqlDB
	return repo.GetUserDetail(db, userId)
}

// create new article
func ArticleNew(conf config.Configuration, article *model.Article) (user model.Article, err error) {

	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	_, err = repo.CreateNewArticle(tx, article)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *article, nil
}

// update article
func ArticleUpdate(conf config.Configuration, article *model.Article) (user model.Article, err error) {

	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	_, err = repo.UpdateArticle(tx, article)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *article, nil
}

// delete article
func UserDelete(conf config.Configuration, article *model.Article) (user model.Article, err error) {
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	_, err = repo.DeleteArticle(tx, article)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *article, nil
}
