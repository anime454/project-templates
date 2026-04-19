package db

import (
	"context"

	"gorm.io/gorm"
)

type Adaptor struct {
	gorm *gorm.DB
}

func NewDB(gorm *gorm.DB) *Adaptor {
	return &Adaptor{
		gorm: gorm,
	}
}

func (a *Adaptor) Close() error {
	sqlDB, err := a.gorm.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (a *Adaptor) Find(dest any, conds ...any) error {
	err := a.gorm.Find(dest, conds...).Error
	if err == gorm.ErrRecordNotFound {
		return ErrRecordNotFound
	}
	return err
}

func (a *Adaptor) Create(value any) error {
	return a.gorm.Create(value).Error
}

func (a *Adaptor) Save(value any) error {
	return a.gorm.Save(value).Error
}

func (a *Adaptor) Delete(value any, conds ...any) error {
	return a.gorm.Delete(value, conds...).Error
}

func (a *Adaptor) First(dest any, conds ...any) error {
	return a.gorm.First(dest, conds...).Error
}

func (a *Adaptor) WithContext(ctx context.Context) *Adaptor {
	return &Adaptor{
		gorm: a.gorm.WithContext(ctx),
	}
}
