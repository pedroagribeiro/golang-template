package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	config "template/config"
	incidenthubLog "template/core/log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var ErrNoRowsDeleted = errors.New("no rows were deleted")

type dbConnectionGorm struct {
	engine *gorm.DB
}

type dbSessionGorm struct {
	session  *gorm.DB
	rollback bool
}

type dbRowGorm struct {
	Session *gorm.DB
	Row     *sql.Rows
}

func (r dbRowGorm) Next(val any) bool {
	if r.Row != nil {
		if r.Row.Next() {
			if err := r.Session.ScanRows(r.Row, val); err != nil {
				incidenthubLog.Errorf("Failed to scan row: %v", err)
				return false
			}
			return true
		}
	}
	return false
}
func (r dbRowGorm) Close() {
	if r.Row != nil {
		r.Row.Close()
	}
}

var gorm_logger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
		Colorful:                  true,
	},
)

var gorm_silent_logger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		LogLevel: logger.Silent,
	},
)

func getLogger() logger.Interface {
	dbLogger := gorm_silent_logger
	cfg := config.GetConfig()
	if !cfg.Database.DbSilent {
		dbLogger = gorm_logger
	}
	return dbLogger
}

func NewDb(sqlConn *sql.DB) (db IDbConnection, err error) {
	// &gorm.Config{Logger: gorm_logger}
	if db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlConn}), &gorm.Config{Logger: getLogger()}); err == nil {
		return &dbConnectionGorm{engine: db}, err
	} else {
		incidenthubLog.Errorf("[NewDb]: Could not initialize database - %s", err.Error())
		return nil, err
	}
}

func (c *dbConnectionGorm) NewSession() (session IDbSession) {
	// output := dbSessionGorm{session: c.engine.Session(&gorm.Session{Logger: gorm_logger})}
	output := dbSessionGorm{session: c.engine.Session(&gorm.Session{Logger: getLogger()})}
	return &output
}

func (c *dbConnectionGorm) Sync(data ...any) (err error) {
	return c.engine.AutoMigrate(data...)
}

func (s *dbSessionGorm) Exec(rawQuery string) error {
	result := s.session.Exec(rawQuery)
	return result.Error
}

func (s *dbSessionGorm) Preload(beans any, conds any, additionalPreloads ...string) error {
	query := s.session.Preload(clause.Associations)
	for _, preload := range additionalPreloads {
		query = query.Preload(preload)
	}
	result := query.First(beans, conds)
	return result.Error
}

func (s *dbSessionGorm) Query(rawQuery string, dest any) error {
	result := s.session.Raw(rawQuery).Scan(&dest)
	return result.Error
}

func (s *dbSessionGorm) QueryArgs(rawQuery string, args []interface{}, dest any) error {
	result := s.session.Raw(rawQuery, args...).Scan(dest)
	return result.Error
}

func (s *dbSessionGorm) QueryRows(values map[string]interface{}, dest any) (value IdbRow, err error) {
	if s.session != nil {
		rows, err := s.session.Where(values).Find(&dest).Rows()
		if err != nil {
			return dbRowGorm{}, err
		}
		return dbRowGorm{Row: rows, Session: s.session}, nil
	}
	return
}

func (s *dbSessionGorm) Create(beans any) error {
	result := s.session.Create(beans)
	return result.Error
}

func (s *dbSessionGorm) CreateWithPreload(beans any) error {
	result := s.session.Create(beans).Preload(clause.Associations)
	return result.Error
}

func (s *dbSessionGorm) CreateInBatch(beans any, batchSize int) error {
	result := s.session.CreateInBatches(beans, batchSize)
	return result.Error
}

func (s *dbSessionGorm) GetJoin(beans any, conds any, join string) (bool, error) {
	result := s.session.Preload(join).Take(beans, conds)

	has := true
	err := result.Error
	if err == gorm.ErrRecordNotFound {
		has, err = false, nil
	}
	return has, err
}

func (s *dbSessionGorm) Get(beans any, conds any) (bool, error) {
	result := s.session.Take(beans, conds)

	has := true
	err := result.Error
	if err != nil && err.Error() == "record not found" {
		has, err = false, nil
	}
	return has, err
}

func (s *dbSessionGorm) FindJoin(beans any, cond any, join string) error {
	result := s.session.Preload(join).Find(beans, cond)
	return result.Error
}

func (s *dbSessionGorm) FindOne(beans any, cond any) error {
	result := s.session.Preload(clause.Associations).Limit(1).Find(beans, cond)
	return result.Error
}

func (s *dbSessionGorm) FindOneWithPreload(beans any, cond any, preloads ...string) error {
	query := s.session.Preload(clause.Associations)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	result := query.Limit(1).Find(beans, cond)
	return result.Error

}

func (s *dbSessionGorm) Find(beans any, cond any) error {
	result := s.session.Preload(clause.Associations).Find(beans, cond)
	return result.Error
}

func (s *dbSessionGorm) FindWithPreload(beans any, cond any, preloads ...string) error {
	query := s.session.Preload(clause.Associations)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	result := query.Find(beans, cond)
	return result.Error
}

func (s *dbSessionGorm) Update(beans any) error {
	result := s.session.Model(beans).Clauses(clause.Returning{}).Updates(beans)
	return result.Error
}

func (s *dbSessionGorm) Save(beans any) error {
	result := s.session.Save(beans)
	return result.Error
}

func (s *dbSessionGorm) Delete(beans any, conds any) error {
	result := s.session.Unscoped().Delete(beans, conds)
	if result.Error != nil {
		return fmt.Errorf("could not delete record: %s", result.Error)
	} else if result.RowsAffected <= 0 {
		return ErrNoRowsDeleted
	}
	return result.Error
}

func (s *dbSessionGorm) DeleteWithConds(beans any, conds any) error {
	result := s.session.Unscoped().Where(conds).Delete(beans)
	if result.Error != nil {
		return fmt.Errorf("could not delete record: %s", result.Error)
	} else if result.RowsAffected <= 0 {
		return ErrNoRowsDeleted
	}
	return result.Error
}

func (s *dbSessionGorm) Model(data any) IDbSession {
	s.session = s.session.Model(data)
	return s
}

func (s *dbSessionGorm) Begin() IDbSession {
	s.session = s.session.Begin()
	s.rollback = true
	return s
}

func (s *dbSessionGorm) Commit() IDbSession {
	s.session = s.session.Commit()
	s.rollback = false
	return s
}

func (s *dbSessionGorm) Rollback() IDbSession {
	s.session = s.session.Rollback()
	s.rollback = false
	return s
}

func (s *dbSessionGorm) RollbackIfNotCommited() IDbSession {
	if s.rollback {
		s.session = s.session.Rollback()
	}
	s.rollback = false
	return s
}

func (s *dbSessionGorm) NativeImpl() *gorm.DB {
	return s.session
}

func (db *Db) GetEngine() IDbConnection {
	return db.engine
}

func (d *Db) SessionGorm() (transaction IDbSession, err error) {
	if d.engine == nil {
		return nil, fmt.Errorf(`DB engine is not initialized`)
	}
	return d.engine.NewSession(), nil
}
