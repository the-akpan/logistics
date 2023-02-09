package models

import "tracka/pkg/config"

func Init() {
	db := config.Get().Db

	initUsers(db)
}
