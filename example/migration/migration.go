package migration

import (
	"github.com/aglis-lab/golem/example/pkg/entity"
	"gorm.io/gorm"
)

func Init(gormDB *gorm.DB) error {
	err := gormDB.AutoMigrate(&entity.User{})
	if err != nil {
		return err
	}

	return nil
}
