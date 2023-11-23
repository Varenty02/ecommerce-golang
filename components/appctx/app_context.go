package appctx

import (
	"gorm.io/gorm"
)

type appContext struct {
	db *gorm.DB
}
type AppContext interface {
	GetMainConnection() *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{
		db: db,
	}
}
func (appCtx *appContext) GetMainConnection() *gorm.DB { return appCtx.db }
