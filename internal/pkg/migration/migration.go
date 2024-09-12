package migration

import (
	"fmt"
	"order-service/config"
	"order-service/internal/pkg/constant"
	"order-service/internal/pkg/logging"

	model "order-service/internal/models/entity"
	dbDriver "order-service/internal/pkg/db"

	"gorm.io/gorm"
)

type Migration interface {
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	if !config.NewEnv().GetBool(constant.MIGRATION_ENABLED) {
		return
	}

	mgConfigurations := map[string]Migration{
		constant.DB_ORDER_SERVICE: &migration{
			DbModels: &[]interface{}{
				model.TOrderModel{},
				model.TOrderLinesModel{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := dbDriver.GetConnection(k)
		if err != nil {
			logging.Log.Error(fmt.Sprintf("Failed to run migration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.AutoMigrate()
			logging.Log.Info(fmt.Sprintf("Successfully run migration for database %s", k))
		}
	}
}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}
