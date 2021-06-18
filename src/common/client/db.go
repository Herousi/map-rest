package client

import (
	"time"

	"github.com/Herousi/map-rest/src/common/conf"
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/names"
	"go.uber.org/zap"
)

var dbCli *xorm.Engine
var pgDbCli *xorm.Engine

func GetDbClient() (xormDB *xorm.Engine, err error) {
	if dbCli != nil {
		return dbCli, nil
	}
	xormDB, err = xorm.NewEngine(conf.Options.DbDriver, conf.Options.DbURL)
	if err != nil {
		zap.L().Error("db connect err", zap.Error(err), zap.String("dbURL", conf.Options.DbURL))
		return nil, err
	}
	if err = xormDB.Ping(); err != nil {
		zap.L().Error("db connect err", zap.Error(err), zap.String("dbURL", conf.Options.DbURL))
		return nil, err
	}
	xormDB.SetMapper(names.SnakeMapper{})
	xormDB.ShowSQL(true)
	xormDB.SetTZLocation(time.UTC)
	xormDB.SetMaxOpenConns(25) // 数据库连接数
	xormDB.SetMaxIdleConns(25) // 空闲数
	xormDB.SetConnMaxLifetime(5 * time.Minute)
	dbCli = xormDB
	zap.L().Info("connect db", zap.String("dbURL", conf.Options.DbURL))
	return dbCli, err
}

func GetPgDbClient() (xormDB *xorm.Engine, err error) {
	if pgDbCli != nil {
		return pgDbCli, nil
	}
	xormDB, err = xorm.NewEngine(conf.Options.PgDbDriver, conf.Options.PgDbURL)
	if err != nil {
		zap.L().Error("db connect err", zap.Error(err), zap.String("pgDbURL", conf.Options.PgDbURL))
		return nil, err
	}
	if err = xormDB.Ping(); err != nil {
		zap.L().Error("db connect err", zap.Error(err), zap.String("pgDbURL", conf.Options.PgDbURL))
		return nil, err
	}
	xormDB.SetMapper(names.SnakeMapper{})
	xormDB.ShowSQL(true)
	xormDB.SetTZLocation(time.UTC)
	xormDB.SetMaxOpenConns(25) // 数据库连接数
	xormDB.SetMaxIdleConns(25) // 空闲数
	xormDB.SetConnMaxLifetime(5 * time.Minute)
	pgDbCli = xormDB
	zap.L().Info("connect db", zap.String("pgDbURL", conf.Options.PgDbURL))
	return pgDbCli, err
}
