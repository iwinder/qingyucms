package db

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/iwinder/qingyucms/app/qycms_comments/internal/conf"
	"github.com/iwinder/qingyucms/app/qycms_comments/internal/data/po"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"sync"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCommentAgentRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	log *log.Helper
}

var (
	mysqlDb *Data
	once    sync.Once
)

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	if strings.EqualFold(conf.Database.Source, "") && mysqlDb.db == nil {
		return &Data{}, cleanup, fmt.Errorf("MySql DB Open failed")
	}
	var err error
	var dbIns *gorm.DB
	l := log.NewHelper(log.With(logger, "module", "mysql/data"))
	once.Do(func() {
		dbIns, err = gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
		mysqlDb = &Data{
			db:  dbIns,
			log: l,
		}
		AutoMigrateTable(dbIns)
	})
	if mysqlDb.db == nil || err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return mysqlDb, cleanup, nil
}

// AutoMigrateTable 初始化table
func AutoMigrateTable(dbIns *gorm.DB) {
	dbIns.AutoMigrate(&po.CommentAgentPO{}, &po.CommentIndexPO{}, &po.CommentContentPO{})
}
