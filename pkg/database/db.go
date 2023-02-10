package database

import (
	"tracka/pkg/config"
)

func Init() {
	db := config.Get().Db

	initUserColl(db)
	initOrders(db)
}
