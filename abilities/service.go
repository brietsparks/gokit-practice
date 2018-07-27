package abilities

import (
	"context"
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"

	"fmt"
	"gokit-practice/util"
)

type Ability struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	OwnerId uuid.UUID `gorm:"type:uuid;primary_key"`
	Caption string
}

type Service interface {
	CreateAbility(ctx context.Context, a Ability) (Ability, error)
	UpdateAbility(ctx context.Context, a Ability) (Ability, error)
	DeleteAbility(ctx context.Context, a Ability) error

	//GetAbility(ctx context.Context, a Ability) (Ability, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) CreateAbility(ctx context.Context, a Ability) (Ability, error) {
	db := s.db.Create(&a)

	if db.Error != nil {
		return a, util.NewUnexpectedError(db.Error)
	}

	return a, nil
}

func (s *service) UpdateAbility(ctx context.Context, a Ability) (Ability, error) {
	db := s.db.Save(&a)

	if db.Error != nil {
		return a, util.NewUnexpectedError(db.Error)
	}

	return a, nil
}

func (s *service) DeleteAbility(ctx context.Context, a Ability) (error) {
	db := s.db.Delete(&a)

	if db.Error != nil {
		return util.NewUnexpectedError(db.Error)
	}

	if db.RowsAffected == 0 {
		return util.NewError(AbilityDNE, "trying to delete non-existent Ability")
	}

	return nil
}

//func (s *service) GetAbility(ctx context.Context, a Ability) (Ability, error) {
//	db := s.db.First(&a)
//
//	return a, db.Error
//}


const (
	AbilityDNE     = "ABILITY_DNE"
)
