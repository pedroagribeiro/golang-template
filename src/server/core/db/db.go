package db

import (
	"database/sql"
	"fmt"
	"strings"
	config "template/config"
	log "template/core/log"
	"time"
)

type IDbConnection interface {
	NewSession() (session IDbSession)
	Sync(...any) (err error)
}

type IDbSession interface {
	Exec(string) (e error)
	Preload(data any, cond any, additionalPreloads ...string) (e error)
	Query(string, any) (e error)
	QueryArgs(string, []interface{}, any) (e error)
	QueryRows(map[string]interface{}, any) (val IdbRow, e error)
	Create(any) (e error)
	CreateWithPreload(any) (e error)
	CreateInBatch(any, int) (e error)
	GetJoin(data any, cond any, join string) (has bool, e error)
	Get(any, any) (has bool, e error)
	FindJoin(data any, cond any, join string) (e error)
	FindOne(data any, cond any) (e error)
	FindOneWithPreload(data any, cond any, preloads ...string) (e error)
	Find(data any, cond any) (e error)
	FindWithPreload(data any, cond any, preloads ...string) (e error)
	Update(data any) (e error)
	Save(data any) (e error)
	Delete(any, any) (e error)
	DeleteWithConds(any, any) error
	Model(any) IDbSession
	Begin() IDbSession
	Commit() IDbSession
	Rollback() IDbSession
	RollbackIfNotCommited() IDbSession
}

type IdbRow interface {
	Next(any) bool
	Close()
}

type DbParams struct {
	Path string `json:"path"`
}

type Db struct {
	engine IDbConnection
}

func createPasswordBit(password string) string {
	if strings.Compare(password, "") != 0 {
		return fmt.Sprintf("password=%s", password)
	}
	return ""
}

func CreateDb(cfg *config.Config) *Db {
	output := Db{}
	output.engine = nil
	log.Infof("[CreateDb]: Database configuration object - %+v", cfg.Database)
	connectionString := fmt.Sprintf("host=%s port=%s user=%s %s dbname=%s sslmode=disable", cfg.Database.DbAddress, cfg.Database.DbPort, cfg.Database.DbUsername, createPasswordBit(cfg.Database.DbPassword), cfg.Database.DbName)
	sqlDB, err := sql.Open("pgx", connectionString)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.DbMaxConnectionsLifetime) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.Database.DbMaxConnectionsIdleTime) * time.Minute)
	sqlDB.SetMaxOpenConns(cfg.Database.DbMaxOpenConnections)
	sqlDB.SetMaxIdleConns(cfg.Database.DbMaxIdleConnections)
	if err != nil {
		log.Errorf("[CreateDb]: Could not initialize DB - %s", err.Error())
	}
	if db, err := NewDb(sqlDB); err != nil {
		log.Errorf("[CreateDb]: Could not initialize DB - %s", err.Error())
	} else {
		output.engine = db
	}
	return &output
}

func (d *Db) Sync(beans ...interface{}) error {
	if d.engine == nil {
		return fmt.Errorf(`DB engine is not initialized`)
	}
	return d.engine.Sync(beans...)
}
